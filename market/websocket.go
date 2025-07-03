package market

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/gorilla/websocket"

	"go.tradeforge.dev/fmp/model"
)

const (
	QuoteEndpoint = "wss://websockets.financialmodelingprep.com"
)

// Connect connects to the FMP websocket endpoint.
//
// NOTE: It is the responsibility of the caller to call Disconnect to close the connection when done.
//
//nolint:gocognit
func (wss *WebsocketClient) Connect(endpoint string) (err error) {
	wss.connectionLock.Lock()
	defer wss.connectionLock.Unlock()

	wss.connectOnce.Do(func() {
		var conn *websocket.Conn
		//nolint:bodyclose // The connection is closed in the Disconnect method.
		conn, _, err = websocket.DefaultDialer.Dial(endpoint, nil)
		if err != nil {
			return
		}
		wss.connection = conn
		wss.manager.Run(wss.maintainConnection)

		msg := model.WebsocketAuthenticationRequest{
			Event: model.WebsocketEventNameLogin,
			Data:  model.WebsocketAuthenticationRequestData{APIKey: wss.config.APIKey},
		}
		if connErr := conn.WriteJSON(msg); connErr != nil {
			err = fmt.Errorf("writing authentication message: %w", connErr)
			return
		}
	L:
		for {
			select {
			case <-wss.ctx.Done():
				break L
			case evt := <-wss.events:
				if evt.Event != model.WebsocketEventNameLogin {
					continue
				}
				if evt.Status == nil || *evt.Status >= 400 {
					errMsg := fmt.Sprintf("unexpected error code: %d", evt.Status)
					if evt.Message != nil {
						errMsg = *evt.Message
					}
					err = fmt.Errorf("authentication failed: %s", errMsg)
				}
				break L
			}
		}
	})
	if err != nil {
		return fmt.Errorf("dialing websocket connection: %w", err)
	}
	return nil
}

func (wss *WebsocketClient) Disconnect() error {
	wss.connectionLock.Lock()
	defer wss.connectionLock.Unlock()
	if wss.connection == nil {
		return nil
	}
	if err := wss.connection.Close(); err != nil {
		return fmt.Errorf("closing websocket connection: %w", err)
	}
	wss.connection = nil
	wss.connectOnce = sync.Once{}

	return nil
}

func (wss *WebsocketClient) Subscribe(symbols []string) error {
	wss.subscribeQuotesLock.Lock()
	defer wss.subscribeQuotesLock.Unlock()
	msg := model.WebsocketSubscriptionRequest{
		Event: model.WebsocketEventNameSubscribe,
		Data:  model.WebsocketSubscriptionRequestData{Symbols: symbols},
	}
	if err := wss.connection.WriteJSON(msg); err != nil {
		return fmt.Errorf("writing subscription message: %w", err)
	}
L:
	for {
		select {
		case <-wss.ctx.Done():
			break L
		case evt := <-wss.events:
			if evt.Event != model.WebsocketEventNameSubscribe {
				continue
			}
			if evt.Status == nil || *evt.Status >= 400 {
				errMsg := fmt.Sprintf("unexpected error code: %d", evt.Status)
				if evt.Message != nil {
					errMsg = *evt.Message
				}
				return fmt.Errorf("subscription failed: %s", errMsg)
			}
			break L
		}
	}
	return nil
}

func (wss *WebsocketClient) Unsubscribe(symbols []string) error {
	wss.subscribeQuotesLock.Lock()
	defer wss.subscribeQuotesLock.Unlock()
	msg := model.WebsocketSubscriptionRequest{
		Event: model.WebsocketEventNameUnsubscribe,
		Data:  model.WebsocketSubscriptionRequestData{Symbols: symbols},
	}
	if err := wss.connection.WriteJSON(msg); err != nil {
		return fmt.Errorf("writing unsubscription message: %w", err)
	}
	return nil
}

func (wss *WebsocketClient) Quotes() <-chan model.WebsocketQuote {
	return wss.quotes
}

//nolint:gocognit
func (wss *WebsocketClient) maintainConnection(ctx context.Context) error {
L:
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			var rawMessage json.RawMessage
			if err := wss.connection.ReadJSON(&rawMessage); err != nil {
				return fmt.Errorf("reading websocket message: %w", err)
			}
			msg := model.WebsocketMesssage{}
			if err := json.Unmarshal(rawMessage, &msg); err != nil {
				return fmt.Errorf("unmarshaling websocket message: %w", err)
			}

			switch msg.Event {
			case model.WebsocketEventNameHeartbeat:
				wss.logger.Debug("received heartbeat")
				continue
			case model.WebsocketEventNameLogin:
				wss.events <- msg
				wss.logger.Debug("authenticated", slog.Any("message", msg))
				continue
			case model.WebsocketEventNameSubscribe:
				wss.events <- msg
				wss.logger.Debug("subscribed", slog.Any("message", msg))
				continue
			case model.WebsocketEventNameUnsubscribe:
				wss.events <- msg
				wss.logger.Debug("unsubscribed", slog.Any("message", msg))
				break L
			default:
				wss.logger.Debug("received message", slog.Any("raw", rawMessage))
				if msg.Type == nil {
					return errors.New("unknown message type: nil")
				}
				if err := wss.processRawMessage(*msg.Type, rawMessage); err != nil {
					return fmt.Errorf("processing message: %w", err)
				}
			}
		}
	}
	return nil
}

func (wss *WebsocketClient) processRawMessage(typ model.WebsocketMessageType, msg json.RawMessage) error {
	switch typ {
	case model.WebsocketMessageTypeQuote:
		if err := wss.processQuote(msg); err != nil {
			return fmt.Errorf("processing quote: %w", err)
		}
	default:
		wss.logger.Debug("received unknown message", slog.Any("message", msg))
	}
	return nil
}

func (wss *WebsocketClient) processQuote(msg json.RawMessage) error {
	quote := model.WebsocketQuote{}
	if err := json.Unmarshal(msg, &quote); err != nil {
		return fmt.Errorf("unmarshaling websocket quote: %w", err)
	}
	wss.quotes <- quote
	return nil
}
