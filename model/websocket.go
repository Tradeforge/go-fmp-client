package model

type WebsocketEventType string

const (
	WebsocketEventTypeHeartbeat   WebsocketEventType = "heartbeat"
	WebsocketEventTypeLogin       WebsocketEventType = "login"
	WebsocketEventTypeSubscribe   WebsocketEventType = "subscribe"
	WebsocketEventTypeUnsubscribe WebsocketEventType = "unsubscribe"
)

type WebsocketEvent struct {
	Event     WebsocketEventType `json:"event"`
	Message   *string            `json:"message"`
	Status    *int               `json:"status"`
	Timestamp *int64             `json:"timestamp"`
}

type WebsocketAuthenticationRequest struct {
	Event WebsocketEventType                 `json:"event"`
	Data  WebsocketAuthenticationRequestData `json:"data"`
}

type WebsocketAuthenticationRequestData struct {
	APIKey string `json:"apiKey" validate:"required"`
}

type WebsocketSubscriptionRequest struct {
	Event WebsocketEventType               `json:"event"`
	Data  WebsocketSubscriptionRequestData `json:"data"`
}

type WebsocketSubscriptionRequestData struct {
	Symbols []string `json:"ticker" validate:"required"`
}

type WebsocketQuote struct {
	Symbol      string  `json:"s"`
	AskPrice    float64 `json:"ap"`
	AskSize     int64   `json:"as"`
	BidPrice    float64 `json:"bp"`
	BidSize     int64   `json:"bs"`
	LastPrice   float64 `json:"lp"`
	LastUpdated int     `json:"t"`
}
