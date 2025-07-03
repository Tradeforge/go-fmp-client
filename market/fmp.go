// Package market defines HTTP and a Websocket client for the FMP API.
package market

import (
	"context"
	"errors"
	"log/slog"
	"sync"

	"github.com/gorilla/websocket"
	"go.tradeforge.dev/background/manager"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

type HTTPClientConfig struct {
	APIKey string `validate:"required" env:"FMP_API_KEY"`
}

// HTTPClient defines a client to the Polygon REST API.
type HTTPClient struct {
	ClockClient
	QuoteClient
	TickerClient
	NewsClient
	EventClient
}

// NewHTTPClient returns a new HTTP client with the specified API key and config.
func NewHTTPClient(
	config HTTPClientConfig,
	logger *slog.Logger,
) *HTTPClient {
	c := rest.New(
		config.APIKey,
		logger,
	)

	return &HTTPClient{
		ClockClient: ClockClient{
			Client: c,
		},
		QuoteClient: QuoteClient{
			Client: c,
		},
		TickerClient: TickerClient{
			Client: c,
		},
		NewsClient: NewsClient{
			Client: c,
		},
		EventClient: EventClient{
			Client: c,
		},
	}
}

type WebsocketClientConfig struct {
	APIKey string `validate:"required" env:"FMP_API_KEY"`
}

type WebsocketClient struct {
	ctx    context.Context
	config WebsocketClientConfig
	logger *slog.Logger

	manager *manager.Manager

	connectOnce    sync.Once
	connectionLock sync.Mutex
	connection     *websocket.Conn

	subscribeQuotesLock sync.RWMutex

	events chan model.WebsocketMesssage
	quotes chan model.WebsocketQuote
}

func NewWebsocketClient(
	ctx context.Context,
	config WebsocketClientConfig,
	logger *slog.Logger,
) (*WebsocketClient, error) {
	if ctx.Done() != nil {
		return nil, errors.New("context is already cancelled")
	}
	return &WebsocketClient{
		ctx:     ctx,
		config:  config,
		logger:  logger,
		manager: manager.New(ctx, manager.WithCancelOnError(), manager.WithFirstError()),

		events: make(chan model.WebsocketMesssage),
		quotes: make(chan model.WebsocketQuote),
	}, nil
}
