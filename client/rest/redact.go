package rest

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/go-resty/resty/v2"
)

// redactedValue is the placeholder written in place of a credential. It
// matches the placeholder the trace logging has always used for the
// Authorization header.
const redactedValue = "REDACTED"

const (
	apiKeyQueryParam  = "apikey"
	apiKeyHeader      = "apikey"
	authorizationName = "authorization"
)

// apiKeyQueryPattern matches the `apikey` query param and its value wherever a
// URL is rendered into text. It stops at the characters that terminate a query
// value in practice — another param, the closing quote net/url puts around a
// URL, or whitespace — so only the credential is replaced.
var apiKeyQueryPattern = regexp.MustCompile(`(?i)(` + apiKeyQueryParam + `=)[^&"'\s]*`)

// redactor removes the FMP credential from text on its way to a log sink or an
// error message. The API is authenticated with an `apikey` query param, so the
// key is in every request URL and therefore in anything that renders one.
type redactor struct {
	apiKey string
}

// redact returns s with the credential replaced by a placeholder, leaving the
// endpoint, the other query params and the underlying cause intact so the
// output still tells an operator which request failed and why.
//
// Two passes, because neither alone is sufficient: the literal pass catches the
// configured key wherever it appears (URL, header, echoed response body), and
// the pattern pass catches any other value of the `apikey` param — a
// per-request override the client never saw.
func (r redactor) redact(s string) string {
	if r.apiKey != "" {
		s = strings.ReplaceAll(s, r.apiKey, redactedValue)
		if escaped := url.QueryEscape(r.apiKey); escaped != r.apiKey {
			s = strings.ReplaceAll(s, escaped, redactedValue)
		}
	}
	return apiKeyQueryPattern.ReplaceAllString(s, "${1}"+redactedValue)
}

// redactHeaders returns a credential-free copy of h. It copies rather than
// rewrites in place: the caller's header map is not ours to mutate.
func (r redactor) redactHeaders(h http.Header) http.Header {
	redacted := make(http.Header, len(h))
	for name, values := range h {
		lowered := strings.ToLower(name)
		cleaned := make([]string, len(values))
		for i, v := range values {
			if lowered == authorizationName || lowered == apiKeyHeader {
				cleaned[i] = redactedValue
				continue
			}
			cleaned[i] = r.redact(v)
		}
		redacted[name] = cleaned
	}
	return redacted
}

// sanitizeError strips the credential from err without flattening it. The
// returned error keeps the original chain underneath, so errors.Is and
// errors.As classify the failure exactly as they did before — a caller can
// still tell a cancellation from a timeout from a refused connection.
func (r redactor) sanitizeError(err error) error {
	if err == nil {
		return nil
	}
	// Taken before the chain is rebuilt, so any context a wrapper added to the
	// message survives.
	message := r.redact(err.Error())

	// A *url.Error keeps the authenticated URL in an exported field of its own.
	// Redacting only the message would still hand the key to a caller doing
	// errors.As, so rebuild it around the same Op and the same cause.
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		err = &url.Error{
			Op:  urlErr.Op,
			URL: r.redact(urlErr.URL),
			Err: urlErr.Err,
		}
	}
	return &redactedError{message: message, err: err}
}

// redactedError carries a credential-free message while keeping the original
// error reachable through Unwrap.
type redactedError struct {
	message string
	err     error
}

func (e *redactedError) Error() string {
	return e.message
}

func (e *redactedError) Unwrap() error {
	return e.err
}

var _ resty.Logger = (*restyLogger)(nil)

// restyLogger adapts resty's logger onto slog and redacts every message on the
// way through.
//
// Resty prints the full request URL on each retry and on the final failure
// (resty/request.go: r.log.Warnf / r.log.Errorf) using a default logger that
// writes straight to stderr. That is the path that put the live API key into
// CloudWatch (BE-475).
type restyLogger struct {
	logger   *slog.Logger
	redactor redactor
}

func (l *restyLogger) Errorf(format string, v ...any) {
	l.logger.Error(l.redactor.redact(fmt.Sprintf(format, v...)))
}

func (l *restyLogger) Warnf(format string, v ...any) {
	l.logger.Warn(l.redactor.redact(fmt.Sprintf(format, v...)))
}

func (l *restyLogger) Debugf(format string, v ...any) {
	l.logger.Debug(l.redactor.redact(fmt.Sprintf(format, v...)))
}
