package rest

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// Every sub-client in this module shares one *rest.Client, and consumers
// paginate several endpoints through it concurrently. CallURL used to call
// c.HTTP.SetTimeout on that shared client for every request, which is a write
// to shared state from every goroutine.
//
// This test only fails under `go test -race`, which is how the bug survived:
// the module's own suite runs without it, and the consumer's `make test` passes
// -race through a variable that is empty locally and set in CI. Green locally,
// red in CI, and the stack trace points into this module rather than at the
// caller.
func TestConcurrentRequestsShareTheClientWithoutRacing(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, `{"ok":true}`)
	}))
	t.Cleanup(srv.Close)

	c := New("TESTKEY_DO_NOT_LOG_concurrency", slog.New(slog.NewTextHandler(io.Discard, nil)))
	c.HTTP.SetBaseURL(srv.URL)

	const goroutines = 8
	var wg sync.WaitGroup
	errs := make([]error, goroutines)

	for i := range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var out map[string]any
			_, err := c.CallURL(context.Background(), http.MethodGet, "/stable/anything", &out)
			errs[i] = err
		}()
	}
	wg.Wait()

	for i, err := range errs {
		require.NoErrorf(t, err, "goroutine %d", i)
	}
}

// The timeout must still be in force after CallURL stopped setting it per
// request -- otherwise removing the racy write would silently drop the timeout
// rather than fix it, and this file would be pinning the absence of a race
// while a 300s hang became unbounded.
func TestClientTimeoutIsSetOnceAtConstruction(t *testing.T) {
	t.Parallel()

	c := New("TESTKEY_DO_NOT_LOG_timeout", slog.New(slog.NewTextHandler(io.Discard, nil)))

	require.Equal(t, DefaultClientTimeout, c.HTTP.GetClient().Timeout)
}
