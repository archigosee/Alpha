# Clinic Suite — Learning Curriculum

Three connected portfolio projects for a fictional clinic, built as a daily learning path.

## Rules
- **2 hours max per day, Monday to Saturday. Sundays are rest days.** Session time is agreed at the end of each session (DZ works mornings on weekdays).
- **Session shape:** Review (10 min) → Concept (20 min) → Build (70 min) → Checkpoint (20 min).
- **The gate:** A day is passed only when DZ answers the checkpoint correctly *in his own words* and can make a small change to the code himself. If not, the next session starts with a redo of the missed part. No new material until the gate is passed.
- **Source of truth:** `PROGRESS.md` in this repo. Every session reads it first and updates it last.
- **Claude writes most of the code; DZ runs every piece locally and explains it back.** Over time, DZ writes more of it himself.
- Day numbers are targets, not deadlines. Redo days push everything back, and that's fine.

## Stack
- Backend: **Go** + **PostgreSQL**
- Frontend: **React** (TypeScript, Vite)
- AI: LLM API + embeddings (pgvector)
- Bot: Telegram
- Deploy: Docker + a simple cloud host

---

## Phase 0 — Foundations (Days 1–6)
| Day | Goal |
|---|---|
| 1 | Tools: Go, Git, VS Code. Repo cloned. First Go program. How HTTP requests and responses work; a "hello" web server. |
| 2 | Go basics: variables, types, functions, errors. Return JSON from an endpoint. |
| 3 | Go structs, slices, maps. Model `Patient`, `Doctor`, `Appointment` in memory. |
| 4 | SQL basics: tables, keys, relationships. Install PostgreSQL. Design the clinic schema on paper first. |
| 5 | SQL queries: SELECT, INSERT, UPDATE, JOIN. Create the schema with migrations. |
| 6 | **Review day:** rebuild Day 1–5 concepts from memory. Phase 0 checkpoint. |

## Phase 1 — Project 1: Booking & Patient Portal
### 1A. Backend API (Days 7–24)
| Days | Goal |
|---|---|
| 7–8 | Project structure (handlers, services, storage layers) and why layering matters. Connect Go to Postgres. |
| 9–11 | CRUD for patients and doctors. Input validation. Proper HTTP status codes. |
| 12–14 | Appointments: availability, booking, preventing double-booking (transactions). |
| 15–17 | Authentication: password hashing, login, sessions/tokens. |
| 18–19 | Authorization: roles (admin, doctor, receptionist, patient). Who can see what. |
| 20–21 | Testing in Go: unit tests and API tests. |
| 22–23 | Config, logging, error handling done properly. API documentation. |
| 24 | **Review day + Phase 1A checkpoint:** DZ draws the whole backend architecture and explains it. |

### 1B. Frontend (Days 25–44)
| Days | Goal |
|---|---|
| 25–27 | JavaScript & TypeScript essentials for React. |
| 28–30 | React: components, props, state. Set up the Vite project. |
| 31–33 | Calling the Go API, loading and error states, forms. |
| 34–36 | Routing, login flow, protected pages by role. |
| 37–39 | Patient portal: book, reschedule, cancel, history. |
| 40–42 | Staff dashboard: daily schedule, patient records, doctor availability. |
| 43 | Amharic/English language switching. |
| 44 | **Review day + Phase 1B checkpoint.** |

### 1C. Ship It (Days 45–50)
| Days | Goal |
|---|---|
| 45–46 | Telegram appointment reminders (background jobs). |
| 47–48 | Docker. Deploy backend, database and frontend. Seed fake demo data. |
| 49 | README with architecture diagram. |
| 50 | Record 2–3 min demo video. **Project 1 shipped.** |

## Phase 2 — Project 2: AI Clinic Assistant (Days 51–70)
| Days | Goal |
|---|---|
| 51–52 | How LLMs work in practice: tokens, prompts, API calls from Go. Cost and limits. |
| 53–54 | Prompt design: system prompts, structured output, refusing medical diagnosis. |
| 55–57 | Embeddings and vector search with pgvector. |
| 58–60 | RAG: answer questions from clinic documents (hours, services, prices, prep instructions). |
| 61–62 | Telegram bot interface. |
| 63–65 | Tool calling: the bot books appointments through Project 1's API. |
| 66 | Amharic support and its limits. Human handoff when unsure. |
| 67–68 | Evaluating answers: a test set of questions, measuring accuracy. |
| 69 | Deploy, README, diagram. |
| 70 | Demo video. **Project 2 shipped.** |

## Phase 3 — Project 3: Security Hardening & Case Study (Days 71–88)
| Days | Goal |
|---|---|
| 71–72 | Security mindset. Threat modeling the clinic suite: what are we protecting, from whom? |
| 73–75 | OWASP Top 10, with each one checked against our own app. |
| 76–78 | Attacking our app with OWASP ZAP: auth, access control, injection. Log every finding. |
| 79–81 | Fixing findings: access checks, rate limiting, security headers, input handling. |
| 82–83 | Protecting patient data: encryption, audit logs (who viewed which record), backups. |
| 84 | AI-specific risks: prompt injection against the bot, and defenses. |
| 85–86 | Build a small Go CLI that scans a website for missing security headers and SSL problems. |
| 87 | Write the before/after security case study. |
| 88 | **Project 3 shipped.** |

## Phase 4 — Portfolio & Launch (Days 89–95)
| Days | Goal |
|---|---|
| 89–90 | Portfolio page linking all three projects, demos and case study. |
| 91–92 | Upwork/Fiverr/LinkedIn profiles positioned around healthcare software. |
| 93–94 | Proposal templates. First outreach. |
| 95 | Final review of the whole journey. |
