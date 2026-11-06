# Progress Log

Every session: read this first, update it last.

## Current status
- **Next session:** Clear redo queue, then Day 5 (SQL queries: SELECT, INSERT, UPDATE, JOIN; create the schema with migrations)
- **Gate status:** Day 4 concepts passed; schema doc fix queued (small, feeds straight into Day 5)
- **Started:** 2026-11-03

**Setup facts:** PostgreSQL **18** runs on port **5433** (old v17 service stopped and disabled, still installed). Database: `alpha_clinic`, user `postgres`. Start-menu SQL Shell is v17's client — harmless version warning; v18's own shell is under the PostgreSQL 18 folder.

**Day 5 review must include (cold):** `len()` returns a count and `==` returns a bool (misread both on Day 4 review); `append` and zero values (missed again on Day 4 review, recovered on re-check).

**Note for tutor:** Some of DZ's written work comes from another AI tool (Day 3 review answer; Day 4 first `docs/schema.md` used SERIAL/VARCHAR/TIMESTAMPTZ/ON DELETE CASCADE and contradicted his own design). Raised kindly twice. When work sounds polished or uses untaught material, ask him to explain it or redo it plainly. His rough own-words answers are usually right. He was tired at the end of Day 4 (late-night session).

## Redo queue
(Concepts DZ missed at a checkpoint. Cleared before any new material.)
- **Fix `docs/schema.md` to match his own Day 4 design:** `date_of_birth` → `DATE` (not TIMESTAMP); `fee` → `NUMERIC(10,2)`; rename `specialist` → `specialty`; FKs must say what they point to (`patient_id INTEGER → patients`, `doctor_id INTEGER → doctors`); add `NOT NULL` to `name` (both tables), `patient_id`, `doctor_id`, `start_at`, `status`; `phone` stays nullable (his decision); status values `booked` / `cancelled` / `completed`. Then Day 5 turns this file into real `CREATE TABLE` SQL.

## Session log
| Date | Day | Covered | Checkpoint | Notes |
|---|---|---|---|---|
| 2026-11-03 | 1 | Installed Go 1.27, Git, VS Code (Windows). DNS vs web server, HTTP request/response, methods, paths, status codes. `go mod init`, `package main`, first Go program. Web server with `net/http`: handlers, catch-all `/`, `/health`, `/about`. | Partial: Q2 (GET/POST with clinic examples) and Q4 (catch-all routing) passed; `/about` route written independently and pushed. Q1 (request/response journey) and Q3 (why the server blocks) missed → redo. | Good instincts; tends to describe what he sees on screen instead of the mechanism behind it. Push him on "who sends what to whom." |
| 2026-11-04 | 1 (redo) | Redo of request/response journey and why `ListenAndServe` blocks. | Passed both from memory. **Day 1 passed.** | Answered from memory while away from his computer. Q3 answer covered the loop but not what happens if the server ended (later requests fail to connect); worth a one-line check on Day 2's review. |
| 2026-11-04 | 2 | Variables, types, static typing, zero values; functions, multiple returns, errors as values, `if err != nil`; JSON and struct tags; `Content-Type` header; query parameters; `http.Error` with 404/400/500; `return` after errors. Built `/doctor` (JSON), `/room` (query param + error), `/fee` (written by DZ). | **Passed.** Checkpoint 4/4. `/fee` task written independently, correct, error check made consistent. Earlier misses (corrected): mistook query param for path; forgot `err` value in a quiz; put calling code at package level; added `error` to a function that can't fail. | Studied concepts away from his computer, then built at home. Asked a good question unprompted (does `room` need declaring in `findRoom`?). Tests things to find answers — encourage. |
| 2026-11-05 | 3 | Review: `switch`, block scope (passed via variation + typed fix), query params with `&`. Structs (type vs value, dot access, zero values), slices (`[]T`, `append`, `len`, indexing from 0, `range`), maps (`map[K]V`, lookup with `found`, unordered). Split code into `models.go` (same package, `go run .`). Built `/patient?id=` (map lookup, `strconv.Atoi`, 400/404), `/patients` (map → slice, `sort.Slice`), `/appointments` (written by DZ, with `?doctor=` filter). | **Passed.** Checkpoint 4/4; follow-ups on his own code (empty result `[]`, `:=` vs `=`) correct. Earlier misses (corrected): `append` without assignment; zero value of unset `Age`; `:=` vs `=` on `err`; slice vs map for ID lookup (recovered at checkpoint). | Session moved to 3pm (work in the morning) for Nov 5–6. Took about 66 minutes. |
| 2026-11-06 | 4 | Review (`:=` vs `=`, `append`, zero values — re-taught). Why databases; tables, rows, columns, SQL types; primary keys; foreign keys and one-to-many; NUMERIC for money; DATE and `YYYY-MM-DD`; phone as TEXT; NOT NULL. Designed clinic schema on paper (two rounds). Installed PostgreSQL 18; fixed forgotten password (two versions running, reset via `pg_hba.conf` trust → reverted to scram-sha-256). `SELECT version()`, `CREATE DATABASE`, `\\l`, `\\c`, `DROP DATABASE`. | **Concepts passed** (checkpoint 5/5, own words; psql tasks done; explained CASCADE/UNIQUE/TIMESTAMPTZ when asked). First `schema.md` was AI-written → asked to rewrite; his rewrite is his own but has errors → redo queue. | Evening session (21:32–23:20), at work during the day. Good persistence through the Postgres install problems. |
