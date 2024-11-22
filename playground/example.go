package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"go.tradeforge.dev/fmp/market"
	"go.tradeforge.dev/fmp/util"
)

const defaultEnvFile = ".env"

func main() {
	ctx := context.Background()
	cfg := util.MustLoadConfig[market.WebsocketClientConfig](defaultEnvFile)
	slog.SetLogLoggerLevel(slog.LevelDebug)
	logger := slog.Default()

	client, err := market.NewWebsocketClient(ctx, cfg, logger)
	if err != nil {
		logger.Error("creating websocket client", slog.Any("error", err))
		os.Exit(1)
	}

	logger.Info("connecting to quotes feed...")
	if err := client.Connect(market.QuoteEndpoint); err != nil {
		logger.Error("connecting to quotes feed", slog.Any("error", err))
		os.Exit(1)
	}
	defer func() {
		if err := client.Disconnect(); err != nil {
			logger.Error("disconnecting", slog.Any("error", err))
		}
	}()
	logger.Info("connected")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	logger.Info("subscribing to quotes feed...")
	if err := client.Subscribe([]string{"AAPL"}); err != nil {
		logger.Error("subscribing to AAPL quotes feed", slog.Any("error", err))
		return
	}
	logger.Info("subscribed")

	for {
		select {
		case <-interrupt:
			logger.Debug("interrupted")
			return
		case quote, ok := <-client.Quotes():
			if !ok {
				logger.Error("quotes feed closed")
				return
			}
			logger.Info("received quote", "quote", quote)
		}
	}
}
