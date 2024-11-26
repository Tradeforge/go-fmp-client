package model

import (
	"encoding/json"
	"log/slog"

	"github.com/shopspring/decimal"
)

type WebsocketEventName string

const (
	WebsocketEventNameHeartbeat   WebsocketEventName = "heartbeat"
	WebsocketEventNameLogin       WebsocketEventName = "login"
	WebsocketEventNameSubscribe   WebsocketEventName = "subscribe"
	WebsocketEventNameUnsubscribe WebsocketEventName = "unsubscribe"
)

type WebsocketMessageType string

const (
	WebsocketMessageTypeQuote WebsocketMessageType = "Q"
)

type WebsocketMesssage struct {
	Event     WebsocketEventName    `json:"event"`
	Type      *WebsocketMessageType `json:"type"`
	Message   *string               `json:"message"`
	Status    *int                  `json:"status"`
	Timestamp *int64                `json:"timestamp"`
}

func (m WebsocketMesssage) LogValue() slog.Value {
	valueMap := map[string]interface{}{
		"event": m.Event,
	}
	if m.Type != nil {
		valueMap["type"] = m.Type
	}
	if m.Message != nil {
		valueMap["message"] = *m.Message
	}
	if m.Status != nil {
		valueMap["status"] = *m.Status
	}
	if m.Timestamp != nil {
		valueMap["timestamp"] = *m.Timestamp
	}
	return slog.AnyValue(valueMap)
}

type WebsocketAuthenticationRequest struct {
	Event WebsocketEventName                 `json:"event"`
	Data  WebsocketAuthenticationRequestData `json:"data"`
}

type WebsocketAuthenticationRequestData struct {
	APIKey string `json:"apiKey" validate:"required"`
}

type WebsocketSubscriptionRequest struct {
	Event WebsocketEventName               `json:"event"`
	Data  WebsocketSubscriptionRequestData `json:"data"`
}

type WebsocketSubscriptionRequestData struct {
	Symbols []string `json:"ticker" validate:"required"`
}

type WebsocketQuote struct {
	Symbol      string          `json:"s"`
	AskPrice    decimal.Decimal `json:"ap"`
	AskSize     decimal.Decimal `json:"as"`
	BidPrice    decimal.Decimal `json:"bp"`
	BidSize     decimal.Decimal `json:"bs"`
	LastPrice   decimal.Decimal `json:"lp"`
	LastUpdated int64           `json:"t"`
}

func (q WebsocketQuote) MarshalBinary() ([]byte, error) {
	return json.Marshal(q)
}
