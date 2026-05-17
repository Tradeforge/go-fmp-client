# TODOS

## Dividends Company endpoint
- **What:** Add `GetDividends(symbol, limit)` for `/stable/dividends` — historical
  dividends for a single company.
- **Why:** Natural companion to the Dividends Calendar; returns the identical record.
- **Reuse:** Reuses `model.GetDividendsCalendarResponse` — no new response struct.
- **Status:** Deferred from the dividends-calendar PR (scoped to calendar only).
