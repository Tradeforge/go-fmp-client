package rest

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRedact(t *testing.T) {
	t.Parallel()

	r := redactor{apiKey: testAPIKey}

	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "the production log line from BE-475",
			in:   `Get "https://financialmodelingprep.com/stable/house-latest?limit=250&page=96&apikey=` + testAPIKey + `": context canceled`,
			want: `Get "https://financialmodelingprep.com/stable/house-latest?limit=250&page=96&apikey=REDACTED": context canceled`,
		},
		{
			name: "the key is not always last",
			in:   "/stable/quote?apikey=" + testAPIKey + "&symbol=AAPL",
			want: "/stable/quote?apikey=REDACTED&symbol=AAPL",
		},
		{
			name: "a value the client never configured",
			in:   "/stable/quote?apikey=some-other-key&symbol=AAPL",
			want: "/stable/quote?apikey=REDACTED&symbol=AAPL",
		},
		{
			name: "the param name is case insensitive",
			in:   "/stable/quote?APIKEY=some-other-key",
			want: "/stable/quote?APIKEY=REDACTED",
		},
		{
			name: "the key outside a URL, e.g. echoed in a body or a header",
			in:   `{"error":"invalid key ` + testAPIKey + `"}`,
			want: `{"error":"invalid key REDACTED"}`,
		},
		{
			name: "an empty value leaves the param recognisable",
			in:   "/stable/quote?apikey=&symbol=AAPL",
			want: "/stable/quote?apikey=REDACTED&symbol=AAPL",
		},
		{
			name: "text with no credential is untouched",
			in:   "/stable/quote?symbol=AAPL",
			want: "/stable/quote?symbol=AAPL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, r.redact(tt.in))
		})
	}
}

func TestRedactWithoutAnAPIKey(t *testing.T) {
	t.Parallel()

	// A client built with no key must not turn every empty string into a
	// placeholder, but the param pattern still applies.
	r := redactor{}

	assert.Equal(t, "/stable/quote?symbol=AAPL", r.redact("/stable/quote?symbol=AAPL"))
	assert.Equal(t, "/stable/quote?apikey=REDACTED", r.redact("/stable/quote?apikey=leaked"))
}

func TestRedactAnEscapedKey(t *testing.T) {
	t.Parallel()

	const awkwardKey = "key/with+chars"
	r := redactor{apiKey: awkwardKey}

	assert.Equal(
		t,
		"/stable/quote?apikey=REDACTED&note=REDACTED",
		r.redact("/stable/quote?apikey="+url.QueryEscape(awkwardKey)+"&note="+awkwardKey),
	)
}

func TestRedactHeadersCopies(t *testing.T) {
	t.Parallel()

	r := redactor{apiKey: testAPIKey}
	headers := http.Header{
		"Authorization": []string{"Bearer " + testAPIKey},
		"Apikey":        []string{testAPIKey},
		"User-Agent":    []string{"Tradeforge client/v0.0.0"},
		"X-Echo":        []string{"/stable/quote?apikey=" + testAPIKey},
	}

	redacted := r.redactHeaders(headers)

	assert.Equal(t, []string{redactedValue}, redacted["Authorization"])
	assert.Equal(t, []string{redactedValue}, redacted["Apikey"])
	assert.Equal(t, []string{"Tradeforge client/v0.0.0"}, redacted["User-Agent"])
	assert.Equal(t, []string{"/stable/quote?apikey=REDACTED"}, redacted["X-Echo"])

	// The caller's map is not ours to rewrite.
	assert.Equal(t, []string{"Bearer " + testAPIKey}, headers["Authorization"])
}

func TestSanitizeErrorKeepsTheChain(t *testing.T) {
	t.Parallel()

	r := redactor{apiKey: testAPIKey}

	require.NoError(t, r.sanitizeError(nil))

	urlErr := &url.Error{
		Op:  "Get",
		URL: "https://financialmodelingprep.com/stable/quote?apikey=" + testAPIKey,
		Err: context.Canceled,
	}
	sanitized := r.sanitizeError(urlErr)

	require.Error(t, sanitized)
	assert.NotContains(t, sanitized.Error(), testAPIKey)
	require.ErrorIs(t, sanitized, context.Canceled)

	var got *url.Error
	require.ErrorAs(t, sanitized, &got)
	assert.Equal(t, "Get", got.Op)
	assert.Equal(t, "https://financialmodelingprep.com/stable/quote?apikey=REDACTED", got.URL)
}

func TestSanitizeErrorKeepsAWrappersMessage(t *testing.T) {
	t.Parallel()

	r := redactor{apiKey: testAPIKey}
	sentinel := errors.New("dial failed")

	wrapped := errors.Join(
		&url.Error{Op: "Get", URL: "https://example.test/?apikey=" + testAPIKey, Err: sentinel},
		errors.New("extra context"),
	)
	sanitized := r.sanitizeError(wrapped)

	require.Error(t, sanitized)
	assert.NotContains(t, sanitized.Error(), testAPIKey)
	assert.Contains(t, sanitized.Error(), "extra context", "wrapper context must survive")
	assert.ErrorIs(t, sanitized, sentinel)
}

func TestSanitizeErrorOnAPlainError(t *testing.T) {
	t.Parallel()

	r := redactor{apiKey: testAPIKey}
	sentinel := errors.New("boom for apikey=" + testAPIKey)

	sanitized := r.sanitizeError(sentinel)

	require.Error(t, sanitized)
	assert.Equal(t, "boom for apikey=REDACTED", sanitized.Error())
	assert.ErrorIs(t, sanitized, sentinel)
}
