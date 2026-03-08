# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go client library for the [Financial Modeling Prep (FMP) API](https://financialmodelingprep.com/developer/docs/). Module path: `go.tradeforge.dev/fmp`. Uses FMP's `/stable/` API endpoints.

## Build & Test Commands

```bash
make fmt           # go fmt ./...
make vet           # go vet ./...
make test          # go test ./... -cover (runs generate + lint first)
make lint          # golangci-lint run (runs generate first)
make generate      # go generate ./...
go test ./pkg/types/ -run TestRangeContains  # run a single test
```

## Architecture

**`market/`** - Domain-specific API clients, each embedding `*rest.Client`:
- `fmp.go` - `HTTPClient` (composes all domain clients) and `WebsocketClient`
- `ticker.go`, `quote.go`, `news.go`, `event.go`, `clock.go`, `disclosure.go`, `analysis.go` - individual domain clients (e.g., `TickerClient`, `QuoteClient`)

**`client/rest/`** - Generic REST client built on `go-resty`. Handles request execution, error parsing, retries, and query param encoding. API key is passed as `apikey` query param.

**`model/`** - Request params, response types, and domain models. Param structs use `path` and `query` struct tags for URL encoding, plus `validate` tags for validation.

**`encoder/`** - URL encoder using `go-playground/form`. Encodes path params (`:param` placeholders) and query params from struct tags.

**`pkg/types/`** - Custom types for JSON deserialization: `Date`, `DateTime`, `TimeHHMM`, `ThousandSeparatedNumeric[T]`, `Range[T]`, `FlexibleBool`, `Empty`.

**`errors/`** - Custom error types.

**`util/`** - Config loader using `caarlos0/env` with `.env` file support and optional `APP_PREFIX`.

## Conventions

- New API endpoints follow the pattern: add path constant + method in `market/<domain>.go`, add param/response structs in `model/<domain>.go`
- Request param structs use `path:"paramName"` for URL path segments and `query:"paramName"` for query parameters
- Many FMP endpoints return arrays; client methods often unwrap the first element (see `GetCompanyProfile`)
- `model.RequestOption` functional options pattern for per-request customization (headers, query params, body, trace, ignored status codes)
- Uses `shopspring/decimal` for financial numeric precision
