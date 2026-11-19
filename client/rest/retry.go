package rest

import (
	"context"
	"crypto/rand"
	"log/slog"
	"math/big"
	"net/http"
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
		if err != nil || !isGatewayStatus(res.StatusCode()) || slices.Contains(ignoredStatuses, res.StatusCode()) || attempt == scheduledAttempts {
			return res, err
		}

		jitter, jitterErr := rand.Int(rand.Reader, big.NewInt(int64(time.Second)))
		if jitterErr != nil {
			return nil, jitterErr
		}
		delay := time.Duration(attempt)*time.Second + time.Duration(jitter.Int64())
		if deadline, ok := ctx.Deadline(); ok && time.Until(deadline) < delay+scheduledAttemptTimeout {
			return res, nil
		}
		c.logger.WarnContext(ctx, "fmp gateway retry", slog.Int("status", res.StatusCode()), slog.Int("attempt", attempt))
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
}
