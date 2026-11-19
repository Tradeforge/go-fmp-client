package rest

import (
	"context"
	"crypto/rand"
	"errors"
	"io"
	"log/slog"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	scheduledAttempts = 3
	// FMP's gateway can take about 60 seconds to send a 504 response.
	scheduledAttemptTimeout = 65 * time.Second
)

func isGatewayStatus(status int) bool {
	return status == http.StatusBadGateway || status == http.StatusServiceUnavailable || status == http.StatusGatewayTimeout
}

func isRetryableTransportError(err error) bool {
	var requestErr *url.Error
	if !errors.As(err, &requestErr) {
		return false
	}
	var networkErr net.Error
	return errors.As(requestErr.Err, &networkErr) || errors.Is(requestErr.Err, io.EOF) || errors.Is(requestErr.Err, io.ErrUnexpectedEOF)
}

func shouldRetryScheduled(res *resty.Response, err error, ignoredStatuses []int) bool {
	if err != nil {
		return isRetryableTransportError(err)
	}
	return isGatewayStatus(res.StatusCode()) && !slices.Contains(ignoredStatuses, res.StatusCode())
}

func (c *Client) executeScheduled(
	ctx context.Context,
	execute func(context.Context) (*resty.Response, error),
	ignoredStatuses []int,
) (*resty.Response, error) {
	for attempt := 1; ; attempt++ {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		attemptCtx, cancel := context.WithTimeout(ctx, scheduledAttemptTimeout)
		res, err := execute(attemptCtx)
		cancel()
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		if !shouldRetryScheduled(res, err, ignoredStatuses) || attempt == scheduledAttempts {
			return res, err
		}

		jitter, jitterErr := rand.Int(rand.Reader, big.NewInt(int64(time.Second)))
		if jitterErr != nil {
			return nil, jitterErr
		}
		delay := time.Duration(attempt)*time.Second + time.Duration(jitter.Int64())
		if deadline, ok := ctx.Deadline(); ok && time.Until(deadline) < delay+scheduledAttemptTimeout {
			return res, err
		}
		if err == nil {
			c.logger.WarnContext(ctx, "fmp gateway retry", slog.Int("status", res.StatusCode()), slog.Int("attempt", attempt))
		} else {
			c.logger.WarnContext(ctx, "fmp transport retry", slog.Int("attempt", attempt))
		}
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
}

// Resty logs an ERROR before Execute returns a transport failure. A scheduled
// attempt cannot own that severity because the next attempt may recover; the
// caller owns terminal transport logging after the bounded loop returns.
type scheduledAttemptLogger struct {
	*restyLogger
	ctx context.Context
}

func (l *scheduledAttemptLogger) Errorf(_ string, _ ...any) {
	l.logger.WarnContext(l.ctx, "fmp request attempt warning")
}

func (l *scheduledAttemptLogger) Warnf(format string, args ...any) {
	l.Errorf(format, args...)
}
