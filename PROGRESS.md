# Progress Log

Every session: read this first, update it last.

## Current status
- **Next session:** Day 4 (SQL basics: tables, keys, relationships; install PostgreSQL; design the clinic schema on paper first)
- **Gate status:** Open — Day 3 passed
- **Started:** 2026-11-03

**Day 4 review must include (cold, from memory):** `:=` vs `=` (missed once on Day 3); why `append` needs `x = append(x, ...)` (missed once); zero values of a struct's unset fields (missed once). Also a quick check that he can explain map vs slice in plain words, not textbook phrasing.

**Note for tutor:** Some of DZ's written answers read like they were copied from another AI tool (polished, textbook phrasing; on Day 3 one answer mentioned "weekend cases" that weren't in the question). This was raised kindly on Day 3. When an answer sounds polished, follow up with a variation or a question about his own code instead of accepting it. His rough, own-words answers are usually correct.

## Redo queue
(Concepts DZ missed at a checkpoint. Cleared before any new material.)
- none

## Session log
| Date | Day | Covered | Checkpoint | Notes |
|---|---|---|---|---|
| 2026-11-03 | 1 | Installed Go 1.27, Git, VS Code (Windows). DNS vs web server, HTTP request/response, methods, paths, status codes. `go mod init`, `package main`, first Go program. Web server with `net/http`: handlers, catch-all `/`, `/health`, `/about`. | Partial: Q2 (GET/POST with clinic examples) and Q4 (catch-all routing) passed; `/about` route written independently and pushed. Q1 (request/response journey) and Q3 (why the server blocks) missed → redo. | Good instincts; tends to describe what he sees on screen instead of the mechanism behind it. Push him on "who sends what to whom." |
| 2026-11-04 | 1 (redo) | Redo of request/response journey and why `ListenAndServe` blocks. | Passed both from memory. **Day 1 passed.** | Answered from memory while away from his computer. Q3 answer covered the loop but not what happens if the server ended (later requests fail to connect); worth a one-line check on Day 2's review. |
| 2026-11-04 | 2 | Variables, types, static typing, zero values; functions, multiple returns, errors as values, `if err != nil`; JSON and struct tags; `Content-Type` header; query parameters; `http.Error` with 404/400/500; `return` after errors. Built `/doctor` (JSON), `/room` (query param + error), `/fee` (written by DZ). | **Passed.** Checkpoint 4/4. `/fee` task written independently, correct, error check made consistent. Earlier misses (corrected): mistook query param for path; forgot `err` value in a quiz; put calling code at package level; added `error` to a function that can't fail. | Studied concepts away from his computer, then built at home. Asked a good question unprompted (does `room` need declaring in `findRoom`?). Tests things to find answers — encourage. |
| 2026-11-05 | 3 | Review: `switch`, block scope (passed via variation + typed fix), query params with `&`. Structs (type vs value, dot access, zero values), slices (`[]T`, `append`, `len`, indexing from 0, `range`), maps (`map[K]V`, lookup with `found`, unordered). Split code into `models.go` (same package, `go run .`). Built `/patient?id=` (map lookup, `strconv.Atoi`, 400/404), `/patients` (map → slice, `sort.Slice`), `/appointments` (written by DZ, with `?doctor=` filter). | **Passed.** Checkpoint 4/4; follow-ups on his own code (empty result `[]`, `:=` vs `=`) correct. Earlier misses (corrected): `append` without assignment; zero value of unset `Age`; `:=` vs `=` on `err`; slice vs map for ID lookup (recovered at checkpoint). | Session moved to 3pm (work in the morning) for Nov 5–6. Took about 66 minutes. |
