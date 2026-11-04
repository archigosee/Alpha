# Progress Log

Every session: read this first, update it last.

## Current status
- **Next session:** Day 2 (Go basics: variables, types, functions, errors; return JSON from an endpoint)
- **Gate status:** Open — Day 1 passed
- **Started:** 2026-11-03

## Redo queue
(Concepts DZ missed at a checkpoint. Cleared before any new material.)
- none

## Session log
| Date | Day | Covered | Checkpoint | Notes |
|---|---|---|---|---|
| 2026-11-03 | 1 | Installed Go 1.27, Git, VS Code (Windows). DNS vs web server, HTTP request/response, methods, paths, status codes. `go mod init`, `package main`, first Go program. Web server with `net/http`: handlers, catch-all `/`, `/health`, `/about`. | Partial: Q2 (GET/POST with clinic examples) and Q4 (catch-all routing) passed; `/about` route written independently and pushed. Q1 (request/response journey) and Q3 (why the server blocks) missed → redo. | Good instincts; tends to describe what he sees on screen instead of the mechanism behind it. Push him on "who sends what to whom." |
| 2026-11-04 | 1 (redo) | Redo of request/response journey and why `ListenAndServe` blocks. | Passed both from memory. **Day 1 passed.** | Answered from memory while away from his computer. Q3 answer covered the loop but not what happens if the server ended (later requests fail to connect); worth a one-line check on Day 2's review. |
