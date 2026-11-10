# Agent instructions

## Shared agent instructions

This `AGENTS.md` is the shared source for Claude Code and Codex. The adjacent `CLAUDE.md` imports it for Claude compatibility. Update this file when changing shared rules. Preserve actual tool names and configuration paths: `.claude/` settings, hooks, Claude model names and reviewer repository names remain accurate even when Codex is doing the work. Use the current host's native tools; Claude-only aliases such as `opus` and `sonnet` are not Codex model IDs. Hooks enforce rules only in hosts where they are registered.

Before editing a subdirectory, read its scoped `AGENTS.md` (or legacy `CLAUDE.md`). When Claude and Codex work concurrently, use a separate worktree and branch for each task and preserve other sessions' changes.

This file provides shared guidance to Claude Code and Codex when working in this repository.

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
