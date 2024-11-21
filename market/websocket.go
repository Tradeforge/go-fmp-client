package market

import (
    "context"
    "fmt"
    "time"

    "github.com/eapache/go-resiliency/retrier"
    "github.com/gorilla/websocket"
    "go.tradeforge.dev/background/manager"

    "go.tradeforge.dev/fmp/model"
)

const (
    defaultConstantBackoffRetries = 5
)

func (wss *WebsocketClient) SubscribeToPriceFeed(ctx context.Context, symbols []string) (<-chan model.WebsocketQuote, error) {
    conn, _, err := websocket.DefaultDialer.Dial(wss.config.APIURL, nil)
    defer func() {
        if err := conn.Close(); err != nil {
            wss.logger.Error("closing websocket connection", "error", err)
        }
    }()
    if err != nil {
        return nil, err
    }

    if err := wss.authenticate(conn); err != nil {
        return nil, fmt.Errorf("authenticating websocket connection: %w", err)
    }
    if err := wss.subscribeToCompanyFeed(conn, symbols); err != nil {
        return nil, fmt.Errorf("subscribing to company feed: %w", err)
    }
    defer func() {
        if err := wss.unsubscribeFromPriceFeed(conn, symbols); err != nil {
            wss.logger.Error("unsubscribing from price feed", "error", err)
        }
    }()
    wss.logger.Debug("subscribed to price feed")

    priceFeed := make(chan model.WebsocketQuote)

    mgr := manager.New(ctx)
    mgr.RunWithRetry(func(ctx context.Context) error {
        defer func() {
            wss.logger.Debug("closing price feed")
            close(priceFeed)
        }()
        return wss.startReadingPriceFeed(ctx, conn, priceFeed)
    }, retrier.New(retrier.ConstantBackoff(defaultConstantBackoffRetries, 5*time.Second), nil))

    return priceFeed, nil
}

func (wss *WebsocketClient) startReadingPriceFeed(ctx context.Context, conn *websocket.Conn, priceFeed chan<- model.WebsocketQuote) error {
L:
    for {
        select {
        case <-ctx.Done():
            return nil
        default:
            var msg model.WebsocketEvent
            if err := conn.ReadJSON(&msg); err != nil {
                return fmt.Errorf("reading websocket message: %w", err)
            }

            switch msg.Event {
            case model.WebsocketEventTypeHeartbeat:
                wss.logger.Debug("received heartbeat")
                continue
            case model.WebsocketEventTypeLogin:
                wss.logger.Debug("authenticated to price feed")
                continue
            case model.WebsocketEventTypeSubscribe:
                wss.logger.Debug("subscribed to price feed")
                continue
            case model.WebsocketEventTypeUnsubscribe:
                wss.logger.Debug("unsubscribed from price feed")
                break L
            default:
                var quote model.WebsocketQuote
                if err := conn.ReadJSON(&quote); err != nil {
                    return fmt.Errorf("reading websocket quote: %w", err)
                }
                if quote.LastPrice == 0 {
                    continue
                }
                priceFeed <- quote
            }
        }
    }

    return nil
}

func (wss *WebsocketClient) authenticate(conn *websocket.Conn) error {
    msg := model.WebsocketAuthenticationRequest{
        Event: model.WebsocketEventTypeLogin,
        Data:  model.WebsocketAuthenticationRequestData{APIKey: wss.config.APIKey},
    }
    if err := conn.WriteJSON(msg); err != nil {
        return fmt.Errorf("writing authentication message: %w", err)
    }
    return nil
}

func (wss *WebsocketClient) subscribeToCompanyFeed(conn *websocket.Conn, symbols []string) error {
    msg := model.WebsocketSubscriptionRequest{
        Event: model.WebsocketEventTypeSubscribe,
        Data:  model.WebsocketSubscriptionRequestData{Symbols: symbols},
    }
    if err := conn.WriteJSON(msg); err != nil {
        return fmt.Errorf("writing subscription message: %w", err)
    }
    return nil
}

func (wss *WebsocketClient) unsubscribeFromPriceFeed(conn *websocket.Conn, symbols []string) error {
    msg := model.WebsocketSubscriptionRequest{
        Event: model.WebsocketEventTypeUnsubscribe,
        Data:  model.WebsocketSubscriptionRequestData{Symbols: symbols},
    }
    if err := conn.WriteJSON(msg); err != nil {
        return fmt.Errorf("writing unsubscription message: %w", err)
    }
    return nil
}
