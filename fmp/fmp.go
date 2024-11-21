// Package fmp defines a REST client for the FMP API.
package fmp

import (
	"log/slog"

	"go.tradeforge.dev/fmp/client"
)

// HTTPClient defines a client to the Polygon REST API.
type HTTPClient struct {
	QuoteClient
	TickerClient
	EventClient
}

// NewHTTPClient returns a new HTTP client with the specified API key and config.
func NewHTTPClient(
	apiURL string,
	apiKey string,
	logger *slog.Logger,
) *HTTPClient {
	c := client.New(
		apiURL,
		apiKey,
		logger,
	)

	return &HTTPClient{
		QuoteClient: QuoteClient{
			Client: c,
		},
		TickerClient: TickerClient{
			Client: c,
		},
		EventClient: EventClient{
			Client: c,
		},
	}
}

type WebsocketClientConfig struct {
	APIKey string `validate:"required"`
	APIURL string `validate:"required"`
}

func NewWebsocketClient(
	config WebsocketClientConfig,
	logger *slog.Logger,
) *WebsocketClient {
	return &WebsocketClient{
		config: config,
		logger: logger,
	}
}

type WebsocketClient struct {
	config WebsocketClientConfig
	logger *slog.Logger
}
