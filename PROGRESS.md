# Progress Log

Every session: read this first, update it last.

## Current status
- **Next session:** Redo (see queue), then Day 2
- **Gate status:** Closed until the redo queue is cleared
- **Started:** 2026-11-03

## Redo queue
(Concepts DZ missed at a checkpoint. Cleared before any new material.)
- **Request/response journey:** DZ must explain, from memory, the full trip of `localhost:8080/health`: the browser sends a GET request with path `/health` → Go's server matches the path to `healthHandler` → the handler writes text into `w` → that text travels back as the response → the browser displays it. (He said "Go sends the request" and "the server responds with the path"; both wrong.)
- **Why `ListenAndServe` blocks:** DZ must explain why the program keeps running: the server is an endless loop waiting for requests; if it ended, no later requests could be answered. Ctrl + C stops it. (He couldn't answer this; it was explained at the end of Day 1.)
- **Small fix:** change the `/about` text to exactly `Alpha Clinic - Addis Ababa. Open Mon-Sat, 8am to 6pm.` and push.

## Session log
| Date | Day | Covered | Checkpoint | Notes |
|---|---|---|---|---|
| 2026-11-03 | 1 | Installed Go 1.27, Git, VS Code (Windows). DNS vs web server, HTTP request/response, methods, paths, status codes. `go mod init`, `package main`, first Go program. Web server with `net/http`: handlers, catch-all `/`, `/health`, `/about`. | Partial: Q2 (GET/POST with clinic examples) and Q4 (catch-all routing) passed; `/about` route written independently and pushed. Q1 (request/response journey) and Q3 (why the server blocks) missed → redo. | Good instincts; tends to describe what he sees on screen instead of the mechanism behind it. Push him on "who sends what to whom." |
