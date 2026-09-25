# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed (breaking)
- `TickerQuote.Timestamp` is now `types.UnixTimestamp` instead of `int64`. FMP
  started serialising the quote timestamp as a fractional number
  (`1768692032.131`), which made every `TickerQuote` response fail to unmarshal.
  The new type accepts both the integral and the fractional form and truncates
  to whole seconds. Call sites that passed the field to `time.Unix(ts, 0)`
  should call `ts.Time()`; call sites that need the raw epoch should call
  `ts.Int64()`.

### Fixed
- `CallURL` no longer sets the timeout on the shared resty client for every
  request (BE-485). Every sub-client hands out the same `*rest.Client`, so a
  consumer issuing concurrent requests raced on that write — detected by
  `go test -race` in `Tradeforge/api`, whose news sync paginates two endpoints
  from two goroutines. The call was redundant as well as unsafe: `New` already
  sets the same value and nothing else mutates it. A per-request deadline
  belongs on the request's context, which `SetContext` already threads through.

### Security
- The FMP API key is no longer written to logs or carried out inside returned errors (BE-475). Two independent paths leaked it. Resty's default logger printed the full authenticated request URL to stderr on every retry and on the final failure, which is how the live key reached CloudWatch. Separately, `CallURL` wrapped the transport `*url.Error` — which holds that same URL in an exported field — straight into the error it returns, so the key propagated into caller logs and Sentry even with the logger fixed. Both now pass through a redactor that replaces the credential with `REDACTED` while leaving the endpoint, the remaining query params and the error chain intact: `errors.Is`/`errors.As` still tell a cancellation from a timeout from a refused connection from a non-2xx status. Upstream response bodies and traced request/response headers are redacted on the same path.
- **The API key must be rotated**, and only now that the redaction is in place — rotating first would put the fresh key straight back into the logs.

## [0.18.0] - 2026-07-27

### Added
- `GetSplitsCalendar` on `EventClient` for `/stable/splits-calendar` — stock split events across all companies for a date range. The response carries `numerator`/`denominator` (a 2:1 forward split is `2`/`1`; a 1:5 reverse split is `1`/`5`) plus a `splitType` discriminator whose observed values are `stock-split` and `stock-dividend`.

## [0.17.0] - 2026-05-22

### Added
- `IndexClient` with `BatchGetIndexQuotes` and `BatchGetIndexShortQuotes` on FMP's `/stable/batch-index-quotes` endpoint; full and short quotes for all market indexes in a single call.
- `IndexClient` is wired into `HTTPClient` and also exposes `GetIndexList` and historical S&P 500 / Nasdaq / Dow Jones constituent endpoints.

## [0.16.0] - 2026-05-17

### Added
- `GetDividendsCalendar` on `EventClient` for `/stable/dividends-calendar` — dividend events across all companies for a date range.

### Fixed
- Resolve pre-existing `golangci-lint` failures.
- Slice-bounds panic in `TestGetEarningsCalendar` on sparse date windows.

### Docs
- Track deferred dividends company endpoint.

## [0.15.3] - 2026-03-21

### Fixed
- `GetAdvancedDCF` now returns multi-year DCF projections instead of collapsing to a single year.

## [0.15.2] - 2026-03-21

### Fixed
- Tolerate non-numeric `amount` values in financial disclosures instead of failing JSON deserialization.

## [0.15.1] - 2026-03-09

### Fixed
- `BatchGetQuotes` (by symbols) returns full `TickerQuote` instead of the short shape.

## [0.15.0] - 2026-03-09

### Changed (breaking)
- Fix the `TickerQuote` model to match the current FMP response shape.

## [0.14.0] - 2026-03-08

### Added
- Test coverage across market clients.

## [0.13.0] - 2026-02-06

### Changed
- Migrate all endpoints to FMP's `/stable/` API surface.

## [0.12.2] - 2026-01-09

### Added
- `EBITPct` field on the advanced DCF analysis response.

## [0.12.1] - 2026-01-09

### Changed
- Update advanced DCF input and response shapes.

## [0.12.0] - 2026-01-09

### Added
- DCF advanced analysis endpoint.

## [0.11.1] - 2025-11-14

### Changed (breaking)
- Rename "latest general news" for consistency.

## [0.11.0] - 2025-11-14

### Changed (breaking)
- Add realtime news per symbol; rename the prior endpoint to "latest".

## [0.10.2] - 2025-11-10

### Changed
- Increase the default client timeout to 300s.

## [0.10.1] - 2025-11-03

### Fixed
- News stable URL.

## [0.10.0] - 2025-11-03

### Changed
- Update news and FMP articles endpoints.

[Unreleased]: https://github.com/Tradeforge/go-fmp-client/compare/v0.18.0...HEAD
[0.18.0]: https://github.com/Tradeforge/go-fmp-client/compare/v0.17.0...v0.18.0
[0.17.0]: https://github.com/Tradeforge/go-fmp-client/compare/v0.16.0...v0.17.0
[0.16.0]: https://github.com/Tradeforge/go-fmp-client/compare/v0.15.3...v0.16.0
[0.15.3]: https://github.com/Tradeforge/go-fmp-client/compare/v0.15.2...v0.15.3
[0.15.2]: https://github.com/Tradeforge/go-fmp-client/compare/v0.15.1...v0.15.2
[0.15.1]: https://github.com/Tradeforge/go-fmp-client/compare/v0.15.0...v0.15.1
[0.15.0]: https://github.com/Tradeforge/go-fmp-client/compare/v0.14.0...v0.15.0
[0.14.0]: https://github.com/Tradeforge/go-fmp-client/compare/v0.13.0...v0.14.0
[0.13.0]: https://github.com/Tradeforge/go-fmp-client/compare/v0.12.2...v0.13.0
[0.12.2]: https://github.com/Tradeforge/go-fmp-client/compare/v0.12.1...v0.12.2
[0.12.1]: https://github.com/Tradeforge/go-fmp-client/compare/v0.12.0...v0.12.1
[0.12.0]: https://github.com/Tradeforge/go-fmp-client/compare/v0.11.1...v0.12.0
[0.11.1]: https://github.com/Tradeforge/go-fmp-client/compare/v0.11.0...v0.11.1
[0.11.0]: https://github.com/Tradeforge/go-fmp-client/compare/v0.10.2...v0.11.0
[0.10.2]: https://github.com/Tradeforge/go-fmp-client/compare/v0.10.1...v0.10.2
[0.10.1]: https://github.com/Tradeforge/go-fmp-client/compare/v0.10.0...v0.10.1
[0.10.0]: https://github.com/Tradeforge/go-fmp-client/releases/tag/v0.10.0
