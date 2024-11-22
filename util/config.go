package util

import (
	"fmt"
	"os"
	"sync"

	"github.com/caarlos0/env/v10"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var (
	once sync.Once

	validate = validator.New()
)

const appPrefixEnvKey = "APP_PREFIX"

func LoadConfig[T any](envFilePath string) (*T, error) {
	loadEnvFile(envFilePath)

	appPrefix := os.Getenv(appPrefixEnvKey)
	if appPrefix != "" {
		appPrefix += "_"
	}
	cfg := new(T)
	if err := env.ParseWithOptions(cfg, env.Options{Prefix: appPrefix}); err != nil {
		return cfg, fmt.Errorf("applying env from file: %w", err)
	}

	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}

	return cfg, nil
}

func MustLoadConfig[T any](envFilePath string) T {
	cfg, err := LoadConfig[T](envFilePath)
	if err != nil {
		panic(fmt.Errorf("loading config: %w", err))
	}
	return *cfg
}

func loadEnvFile(path string) {
	once.Do(func() {
		if path == "" {
			path = ".env"
		}

		if err := godotenv.Load(path); err != nil {
			if !os.IsNotExist(err) {
				panic(fmt.Errorf("loading env file: %w", err))
			}
		}
	})
}
