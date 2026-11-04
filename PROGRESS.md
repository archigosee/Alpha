# Progress Log

Every session: read this first, update it last.

## Current status
- **Next session:** Day 3 (Go structs, slices, maps; model Patient, Doctor, Appointment in memory)
- **Gate status:** Open — Day 2 passed
- **Started:** 2026-11-03

**Day 3 review must include (cold, from memory):** explain `switch` and why `var isSpecialist bool` is declared before it (block scope). DZ used `switch` on his own on Day 2 before it was taught; his written explanation was correct but polished, so confirm he can say it in his own words. Also: query parameter vs path (missed once on Day 2).

## Redo queue
(Concepts DZ missed at a checkpoint. Cleared before any new material.)
- none

## Session log
| Date | Day | Covered | Checkpoint | Notes |
|---|---|---|---|---|
| 2026-11-03 | 1 | Installed Go 1.27, Git, VS Code (Windows). DNS vs web server, HTTP request/response, methods, paths, status codes. `go mod init`, `package main`, first Go program. Web server with `net/http`: handlers, catch-all `/`, `/health`, `/about`. | Partial: Q2 (GET/POST with clinic examples) and Q4 (catch-all routing) passed; `/about` route written independently and pushed. Q1 (request/response journey) and Q3 (why the server blocks) missed → redo. | Good instincts; tends to describe what he sees on screen instead of the mechanism behind it. Push him on "who sends what to whom." |
| 2026-11-04 | 1 (redo) | Redo of request/response journey and why `ListenAndServe` blocks. | Passed both from memory. **Day 1 passed.** | Answered from memory while away from his computer. Q3 answer covered the loop but not what happens if the server ended (later requests fail to connect); worth a one-line check on Day 2's review. |
| 2026-11-04 | 2 | Variables, types, static typing, zero values; functions, multiple returns, errors as values, `if err != nil`; JSON and struct tags; `Content-Type` header; query parameters; `http.Error` with 404/400/500; `return` after errors. Built `/doctor` (JSON), `/room` (query param + error), `/fee` (written by DZ). | **Passed.** Checkpoint 4/4. `/fee` task written independently, correct, error check made consistent. Earlier misses (corrected): mistook query param for path; forgot `err` value in a quiz; put calling code at package level; added `error` to a function that can't fail. | Studied concepts away from his computer, then built at home. Asked a good question unprompted (does `room` need declaring in `findRoom`?). Tests things to find answers — encourage. |
