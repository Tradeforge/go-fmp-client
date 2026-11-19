# BE-1194 Scheduled FMP Retry Implementation Plan

Doc: https://chatgpt.com/space/page_d5e22cf527c881919ec88cba73d9791d

> Implement in the existing isolated worktree; run affected tests and full CI.

**Goal:** recover scheduled transient gateway failures without paging on intermediate attempts.
**Spec:** [approved design](../specs/2026-11-19-be-1194-scheduled-retries-design.md)
**Architecture:** a typed context marker opts GETs into a request-local retry loop. Keep shared Resty immutable during requests and return existing ResponseError.
**Tech:** Go,Resty2.13,slog,httptest,testify.

## Global constraints
Three attempts maximum; statuses502/503/504 only; delays[1,2)s then[2,3)s;65s per-attempt cap; remaining deadline must fit wait plus65s before retry. Non-GET and unmarked context unchanged. No raw response/body/query in new warning logs. Scope only go-fmp-client; API and infra changes coordinated in separate repositories.

## Task1: bounded gateway recovery
- [ ] Add model.WithScheduledRetries(context.Context) context.Context and IsScheduledRetryContext(context.Context) bool in model/request.go using private key.
- [ ] Add client/rest/retry_test.go first; prove current implementation cannot recover502→success. Table test502/503/504, exhaustion exactly3, default and POST one gateway attempt,404/429/500 no retry, ignored gateway status unchanged, cancellation and insufficient deadline, and logs never include sentinel body.
- [ ] Implement scheduled execution in client/rest/retry.go: request-local AddRetryCondition(false), synchronous loop, context65s attempt cap, exponential additive jitter and cancellable waits. Keep public method signatures and default flow.
- [ ] Handle terminal scheduled gateway at WARN rather than ERROR and retain structured ResponseError; other statuses retain ERROR. Fresh error/result handling must not carry an earlier failed response into success.
- [ ] Verify red then green with go test ./client/rest ./model; go test -race ./... (core HTTP code affects all consumers), go vet ./..., make lint. Capture exact tree and logs in /tmp/be1194-fmp-*.log.
- [ ] Use commit skill, ship-lean in fresh subagent, inspect full current-head CI. No review workflow exists on this public module; report rather than pretending reviewer-bot ran. Root owns reviewed merge/tag.

## Review focus
Concurrent marked/unmarked calls; deadline skipped retry returns typed gateway error; cancellation never retried; ignored status no warning; HTTP response body absent from warnings.

## Engineering review
Scope accepted as-is: six estimated files, no new service. Existing request options and ResponseError reused. Two competing shapes inspected; local loop chosen because Resty2.13 RetryCount/Wait are client fields and Clone is explicitly unsafe concurrently (dependency client.go1137). No user-visible contract change.
Architecture: no issue after loop avoids global mutation; model context helper is the only public addition. Code quality: keep one loop, no generic retry framework. Tests: existing classification/redaction/concurrency tests preserved; new behavior exercised via local HTTP server. Performance: <200s sequential worst-case request budget, bounded by caller; no work on unrelated requests. Outside voice unavailable under this Codex harness; native source review completed, missing independent external coverage recorded.

GET -> marked? no -> existing Resty -> existing classification
             yes -> attempt(max65s) -> success/other status -> return
                                  -> gateway -> ignored? return
                                             -> remaining tries+deadline? warn,wait,next
                                             -> exhausted/budget -> terminal WARN,typed error
                                  -> transport/cancel -> existing transport error return

No tests retired. No additional TODO. API uptake follows library tag; root coordinates3 PRs(library,API,infra). Full CI workflow already runs race/coverage across ./....

## GSTACK REVIEW REPORT

| Review | Trigger | Why | Runs | Status | Findings |
|---|---|---|---|---|---|
| Eng Review | /plan-eng-review | Architecture and tests |1| CLEAR |0 issues; no critical gaps|
| Outside Review | Codex harness preflight | Independent opinion |0| unavailable |No external review; native review only|

**VERDICT:** ENG CLEARED; required bounds and tests mapped.
NO UNRESOLVED DECISIONS
