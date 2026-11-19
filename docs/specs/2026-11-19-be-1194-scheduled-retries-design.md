# BE-1194: scheduled FMP gateway recovery

Doc: https://chatgpt.com/space/page_e4eb4547c1a88191a3d1a1af6bbd9634

Owner approved 2026-11-19. Plan: [implementation](../plans/2026-11-19-be-1194-scheduled-retries.md).

Scheduled GET calls marked with model.WithScheduledRetries(ctx) retry HTTP 502,503,504 twice, for three total attempts. Backoff uses jitter in [1s,2s) then [2s,3s). Each attempt has a65s context timeout. A retry starts only if its wait plus a full attempt fits the parent deadline. Cancellation stops waiting and requests. The API opts in only its two affected scheduled handlers; interactive calls preserve current retry/timeout behavior.

Intermediate gateway failures log a structured warning without response body or request text. Exhausted known scheduled gateway failures also warn once at the library boundary and return model.ResponseError to the caller; Lambda failure metrics, sustained alarm, heartbeat and DLQ retain escalation. Unexpected statuses preserve error logging. This avoids one lower-layer ERROR bypassing the agreed scheduled alert policy.

Approaches: a shared Resty client mutation was rejected as racy; cloning Resty per request was rejected because Clone explicitly disallows concurrent use. Use a small request-local bounded retry loop, disabling Resty's nested retry for opted-in GETs only. Existing transport retries remain unchanged outside the opt-in; inside, transport failures are returned immediately because this policy covers HTTP gateway responses only.

Non-goals: global interactive latency changes,429 policy,other scheduled job adoption,API schemas,DLQ purge,market-error rewrite. Library release precedes API module bump. Worst case three65s attempts plus under5s waits; caller deadline always wins. Active estimate45–75min; CI/review30–60min.
