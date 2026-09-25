package rest

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/model"
)

// testAPIKey is a sentinel: it must never appear in a log line or in a
// returned error, and it is distinctive enough that a substring assertion
// cannot match it by accident.
const testAPIKey = "TESTKEY_DO_NOT_LOG_abc123"

// testPath mirrors the endpoint from BE-475's production log line, query
// params included, so the redacted output can be checked for usefulness.
const testPath = "/stable/house-latest?limit=250&page=96"

const (
	testRetryWaitTime    = time.Millisecond
	testRetryMaxWaitTime = time.Millisecond
	testRequestTimeout   = 250 * time.Millisecond
)

// newTestClient returns a client wired to a capturing logger. Retry waits are
// squashed so the three default retries cost nothing.
func newTestClient(t *testing.T, baseURL string) (*Client, *bytes.Buffer) {
	t.Helper()

	buf := &bytes.Buffer{}
	logger := slog.New(slog.NewTextHandler(buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	c := New(testAPIKey, logger)
	c.HTTP.SetBaseURL(baseURL)
	c.HTTP.SetRetryWaitTime(testRetryWaitTime)
	c.HTTP.SetRetryMaxWaitTime(testRetryMaxWaitTime)

	return c, buf
}

// unreachableBaseURL binds a port and immediately releases it, so every
// request to it fails at the transport layer.
func unreachableBaseURL(t *testing.T) string {
	t.Helper()

	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	addr := l.Addr().String()
	require.NoError(t, l.Close())

	return "http://" + addr
}

func TestTransportFailureIsNotLoggedWithTheAPIKey(t *testing.T) {
	t.Parallel()

	c, logs := newTestClient(t, unreachableBaseURL(t))

	var response map[string]any
	_, err := c.CallURL(context.Background(), http.MethodGet, testPath, &response)
	require.Error(t, err)

	assert.NotContains(t, logs.String(), testAPIKey, "resty logged the API key")
	assert.NotEmpty(t, logs.String(), "the failure must still be logged")
}

func TestTransportFailureIsNotReturnedWithTheAPIKey(t *testing.T) {
	t.Parallel()

	c, _ := newTestClient(t, unreachableBaseURL(t))

	var response map[string]any
	_, err := c.CallURL(context.Background(), http.MethodGet, testPath, &response)
	require.Error(t, err)

	assert.NotContains(t, err.Error(), testAPIKey, "the returned error carries the API key")

	// A caller that reaches for the underlying *url.Error must not find the
	// key in its URL field either.
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		assert.NotContains(t, urlErr.URL, testAPIKey, "url.Error.URL carries the API key")
	}
}

func TestRedactedOutputStillIdentifiesTheRequest(t *testing.T) {
	t.Parallel()

	c, logs := newTestClient(t, unreachableBaseURL(t))

	var response map[string]any
	_, err := c.CallURL(context.Background(), http.MethodGet, testPath, &response)
	require.Error(t, err)

	for _, want := range []string{"/stable/house-latest", "limit=250", "page=96", "apikey=REDACTED"} {
		assert.Contains(t, logs.String(), want, "redaction destroyed information an operator needs")
		assert.Contains(t, err.Error(), want, "redaction destroyed information a caller needs")
	}
}

func TestErrorClassificationSurvivesRedaction(t *testing.T) {
	t.Parallel()

	t.Run("context cancellation", func(t *testing.T) {
		t.Parallel()

		c, _ := newTestClient(t, unreachableBaseURL(t))
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		var response map[string]any
		_, err := c.CallURL(ctx, http.MethodGet, testPath, &response)
		require.Error(t, err)
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("deadline", func(t *testing.T) {
		t.Parallel()

		srv := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
			<-r.Context().Done()
		}))
		defer srv.Close()

		c, _ := newTestClient(t, srv.URL)
		ctx, cancel := context.WithTimeout(context.Background(), testRequestTimeout)
		defer cancel()

		var response map[string]any
		_, err := c.CallURL(ctx, http.MethodGet, testPath, &response)
		require.Error(t, err)
		require.ErrorIs(t, err, context.DeadlineExceeded)

		var netErr net.Error
		require.ErrorAs(t, err, &netErr)
		assert.True(t, netErr.Timeout(), "a timeout must still report itself as one")
	})

	t.Run("refused connection", func(t *testing.T) {
		t.Parallel()

		c, _ := newTestClient(t, unreachableBaseURL(t))

		var response map[string]any
		_, err := c.CallURL(context.Background(), http.MethodGet, testPath, &response)
		require.Error(t, err)

		var urlErr *url.Error
		require.ErrorAs(t, err, &urlErr)
		// net/url renders the verb in title case; the point is that Op is
		// preserved rather than flattened away.
		assert.Equal(t, "Get", urlErr.Op)

		var opErr *net.OpError
		assert.ErrorAs(t, err, &opErr, "the dial failure must stay reachable")
	})

	t.Run("response status", func(t *testing.T) {
		t.Parallel()

		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(`{"error":"forbidden"}`))
		}))
		defer srv.Close()

		c, _ := newTestClient(t, srv.URL)

		var response map[string]any
		res, err := c.CallURL(context.Background(), http.MethodGet, testPath, &response)
		require.Error(t, err)
		require.NotNil(t, res)
		assert.Equal(t, http.StatusForbidden, res.StatusCode())

		var responseErr *model.ResponseError
		require.ErrorAs(t, err, &responseErr)
		assert.Equal(t, http.StatusForbidden, responseErr.StatusCode)
	})
}

func TestResponseErrorEchoingTheURLIsRedacted(t *testing.T) {
	t.Parallel()

	// Some upstreams quote the request URL back in the error body. If they do,
	// the key must not ride out through the error or the log.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error":"invalid key for ` + r.URL.String() + `"}`))
	}))
	defer srv.Close()

	c, logs := newTestClient(t, srv.URL)

	var response map[string]any
	_, err := c.CallURL(context.Background(), http.MethodGet, testPath, &response)
	require.Error(t, err)

	assert.NotContains(t, logs.String(), testAPIKey, "an echoed URL leaked the key into the log")
	assert.NotContains(t, err.Error(), testAPIKey)

	var responseErr *model.ResponseError
	require.ErrorAs(t, err, &responseErr)
	assert.NotContains(t, responseErr.ErrorMessage, testAPIKey, "ResponseError.ErrorMessage carries the key")
}

func TestTraceDoesNotLogCredentials(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	c, logs := newTestClient(t, srv.URL)

	var response map[string]any
	_, err := c.CallURL(
		context.Background(),
		http.MethodGet,
		testPath,
		&response,
		model.WithTrace(true),
		model.Header("Authorization", "Bearer "+testAPIKey),
		model.Header("Apikey", testAPIKey),
	)
	require.NoError(t, err)

	assert.NotContains(t, logs.String(), testAPIKey, "trace logging leaked a credential header")
	assert.Contains(t, logs.String(), "REDACTED")
}

// A per-request override of the apikey query param carries a value the client
// has never seen, so redaction cannot rely on knowing the key.
func TestUnknownAPIKeyValueIsStillRedacted(t *testing.T) {
	t.Parallel()

	const overrideKey = "OVERRIDEKEY_DO_NOT_LOG_xyz789"

	c, logs := newTestClient(t, unreachableBaseURL(t))

	var response map[string]any
	_, err := c.CallURL(
		context.Background(),
		http.MethodGet,
		"/stable/house-latest",
		&response,
		model.QueryParam("apikey", overrideKey),
	)
	require.Error(t, err)

	assert.NotContains(t, logs.String(), overrideKey)
	assert.NotContains(t, err.Error(), overrideKey)
}
