package market

import (
	"log/slog"
	"os"
	"testing"
)

func newTestHTTPClient(t *testing.T) *HTTPClient {
	t.Helper()

	apiKey := os.Getenv("FMP_API_KEY")
	if apiKey == "" {
		t.Skip("FMP_API_KEY not set, skipping integration test")
	}

	return NewHTTPClient(
		HTTPClientConfig{APIKey: apiKey},
		slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})),
	)
}
