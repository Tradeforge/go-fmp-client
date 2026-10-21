package market

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

// BE-777: GetSenateFinancialDisclosures called the House path, so every Senate
// sync fetched House disclosures. The integration tests above could not notice:
// both endpoints return rows of the same shape. This one runs offline and pins
// the path each method requests.
func TestFinancialDisclosuresRequestTheirOwnPath(t *testing.T) {
	var (
		mu    sync.Mutex
		paths []string
	)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		paths = append(paths, r.URL.Path)
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("[]"))
	}))
	t.Cleanup(srv.Close)

	c := rest.New("test-key", slog.New(slog.NewTextHandler(io.Discard, nil)))
	c.HTTP.SetBaseURL(srv.URL)
	c.HTTP.SetRetryCount(0)
	dc := DisclosureClient{Client: c}
	ctx := context.Background()

	_, err := dc.GetHouseFinancialDisclosures(ctx, model.GetHouseFinancialDisclosuresParams{})
	require.NoError(t, err)
	_, err = dc.GetSenateFinancialDisclosures(ctx, model.GetSenateFinancialDisclosuresParams{})
	require.NoError(t, err)

	assert.Equal(t, []string{GetHouseFinancialDisclosuresPath, GetSenateFinancialDisclosuresPath}, paths)
	assert.Equal(t, "/stable/senate-latest", GetSenateFinancialDisclosuresPath)
}
