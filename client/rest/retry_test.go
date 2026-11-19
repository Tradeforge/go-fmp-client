package rest

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/model"
)

func TestScheduledGatewayRecovery(t *testing.T) {
	t.Parallel()
	for _, status := range []int{502, 503, 504} {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			t.Parallel()
			var attempts atomic.Int32
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if attempts.Add(1) == 1 {
					w.WriteHeader(status)
					_, _ = w.Write([]byte(`{"error":"PRIVATE_BODY_SENTINEL"}`))
					return
				}
				_, _ = w.Write([]byte(`{"ok":true}`))
			}))
			t.Cleanup(srv.Close)
			c, logs := newTestClient(t, srv.URL)
			var out map[string]bool
			started := time.Now()
			_, err := c.CallURL(model.WithScheduledRetries(context.Background()), http.MethodGet, "/stable/test", &out)
			require.NoError(t, err)
			assert.True(t, out["ok"])
			assert.EqualValues(t, 2, attempts.Load())
			assert.GreaterOrEqual(t, time.Since(started), time.Second)
			assert.Contains(t, logs.String(), "fmp gateway retry")
			assert.NotContains(t, logs.String(), "level=ERROR")
			assert.NotContains(t, logs.String(), "PRIVATE_BODY_SENTINEL")
		})
	}
}

func TestScheduledGatewayRetryBoundaries(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name      string
		status    int
		method    string
		scheduled bool
		ignored   bool
		attempts  int32
	}{
		{"exhausted", 502, http.MethodGet, true, false, 3},
		{"interactive", 502, http.MethodGet, false, false, 1},
		{"post", 502, http.MethodPost, true, false, 1},
		{"not-found", 404, http.MethodGet, true, false, 1},
		{"rate-limited", 429, http.MethodGet, true, false, 1},
		{"internal", 500, http.MethodGet, true, false, 1},
		{"ignored", 502, http.MethodGet, true, true, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var attempts atomic.Int32
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				attempts.Add(1)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(`{"error":"PRIVATE_BODY_SENTINEL"}`))
			}))
			t.Cleanup(srv.Close)
			c, logs := newTestClient(t, srv.URL)
			ctx := context.Background()
			if tc.scheduled {
				ctx = model.WithScheduledRetries(ctx)
			}
			var opts []model.RequestOption
			if tc.ignored {
				opts = append(opts, model.WithIgnoredErrorStatusCodes(tc.status))
			}
			var out map[string]bool
			_, err := c.CallURL(ctx, tc.method, "/stable/test", &out, opts...)
			assert.Equal(t, tc.attempts, attempts.Load())
			if tc.ignored {
				require.NoError(t, err)
				assert.Empty(t, logs.String())
				return
			}
			var responseErr *model.ResponseError
			require.ErrorAs(t, err, &responseErr)
			assert.Equal(t, tc.status, responseErr.StatusCode)
			if tc.name == "exhausted" {
				assert.Equal(t, 2, strings.Count(logs.String(), "fmp gateway retry"))
				assert.Equal(t, 1, strings.Count(logs.String(), "fmp gateway failure"))
				assert.NotContains(t, logs.String(), "level=ERROR")
				assert.NotContains(t, logs.String(), "PRIVATE_BODY_SENTINEL")
				assert.Contains(t, logs.String(), "fmp gateway failure")
			} else {
				assert.Contains(t, logs.String(), "level=ERROR")
			}
		})
	}
}

func TestScheduledGatewayDeadlineSkipsRetry(t *testing.T) {
	t.Parallel()
	var attempts atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusGatewayTimeout)
	}))
	t.Cleanup(srv.Close)
	c, logs := newTestClient(t, srv.URL)
	ctx, cancel := context.WithTimeout(model.WithScheduledRetries(context.Background()), time.Second)
	defer cancel()
	_, err := c.CallURL(ctx, http.MethodGet, "/stable/test", nil)
	var responseErr *model.ResponseError
	require.ErrorAs(t, err, &responseErr)
	assert.EqualValues(t, 1, attempts.Load())
	assert.NotContains(t, logs.String(), "level=ERROR")
}

func TestScheduledGatewayPreCanceledDoesNotSend(t *testing.T) {
	t.Parallel()
	var attempts atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusBadGateway)
	}))
	t.Cleanup(srv.Close)
	c, _ := newTestClient(t, srv.URL)
	ctx, cancel := context.WithCancel(model.WithScheduledRetries(context.Background()))
	cancel()
	_, err := c.CallURL(ctx, http.MethodGet, "/stable/test", nil)
	require.ErrorIs(t, err, context.Canceled)
	assert.Zero(t, attempts.Load())
}

func TestScheduledGatewayCancelDuringBackoff(t *testing.T) {
	t.Parallel()
	var attempts atomic.Int32
	firstResponse := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if attempts.Add(1) == 1 {
			close(firstResponse)
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	t.Cleanup(srv.Close)
	c, _ := newTestClient(t, srv.URL)
	ctx, cancel := context.WithCancel(model.WithScheduledRetries(context.Background()))
	t.Cleanup(cancel)
	done := make(chan error, 1)
	go func() {
		_, err := c.CallURL(ctx, http.MethodGet, "/stable/test", nil)
		done <- err
	}()
	<-firstResponse
	// Allow the response to enter the retry wait, then cancel before its1s floor.
	time.AfterFunc(100*time.Millisecond, cancel)
	select {
	case err := <-done:
		require.ErrorIs(t, err, context.Canceled)
	case <-time.After(time.Second):
		t.Fatal("cancellation did not interrupt retry wait")
	}
	assert.EqualValues(t, 1, attempts.Load())
}

func TestConcurrentScheduledAndInteractiveRequests(t *testing.T) {
	t.Parallel()
	var scheduled, interactive atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/scheduled" && scheduled.Add(1) == 1 {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		if r.URL.Path == "/interactive" {
			interactive.Add(1)
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	t.Cleanup(srv.Close)
	c, _ := newTestClient(t, srv.URL)
	// A discard handler is concurrency-safe; bytes.Buffer is intentionally not.
	c.logger = slog.New(slog.NewTextHandler(io.Discard, nil))
	c.HTTP.SetLogger(&restyLogger{logger: c.logger, redactor: c.redactor})
	var wg sync.WaitGroup
	errs := make([]error, 2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, errs[0] = c.CallURL(model.WithScheduledRetries(context.Background()), http.MethodGet, "/scheduled", nil)
	}()
	go func() {
		defer wg.Done()
		_, errs[1] = c.CallURL(context.Background(), http.MethodGet, "/interactive", nil)
	}()
	wg.Wait()
	require.NoError(t, errs[0])
	require.Error(t, errs[1])
	assert.EqualValues(t, 2, scheduled.Load())
	assert.EqualValues(t, 1, interactive.Load())
}

type failingRoundTripper struct{ attempts atomic.Int32 }

func (f *failingRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	f.attempts.Add(1)
	return nil, errors.New("transport failure")
}

func TestScheduledTransportFailureDoesNotMultiplyRetries(t *testing.T) {
	t.Parallel()
	c, logs := newTestClient(t, "https://example.invalid")
	transport := &failingRoundTripper{}
	c.HTTP.SetTransport(transport)
	_, err := c.CallURL(model.WithScheduledRetries(context.Background()), http.MethodGet, "/stable/test", nil)
	require.Error(t, err)
	assert.EqualValues(t, 1, transport.attempts.Load())
	assert.Contains(t, logs.String(), "level=ERROR")
	assert.NotContains(t, logs.String(), testAPIKey)
}
