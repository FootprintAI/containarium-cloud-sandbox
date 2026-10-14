# QA report (DRAFT): ReadingTime helper

**Date:** 2026-10-15
**Status:** draft — not executed
**Owner:** unassigned
**Flow:** f16c5d49, stage 5 of 6 (qa)

## What this report is

A draft checklist against the sprint's seven acceptance criteria (see
`docs/qa/reading-time-test-plan.md` for how each one is checked). **No
test was executed on this box in this session, and every result below
is still to be filled in by a person.** Nothing in this report should be
read as a pass.

## Why this is a documentation pass, not an environment run

`docs/sprint/reading-time.md` (the sprint stage's own output) recorded
this sprint's verification queue as empty: every acceptance criterion is
pure logic, provable by `go vet ./... && go test ./...`, and none of it
needs a deployed environment or a UI flow. There is no umbrella
verification queue item to batch here, no deploy record, and no running
service to confirm a commit against — the preconditions this skill
normally requires before an environment pass simply don't apply to this
piece of work. What follows is instead a checklist mapping each
criterion to the code/tests that should satisfy it, for a person to
confirm against an actual test run.

## What was actually done in this session

- Read `text/reading_time.go`, `text/reading_time_test.go`, and
  `README.md` as merged onto this flow branch.
- Mapped each of the issue's seven acceptance criteria to the specific
  test row(s) or code that should satisfy it (recorded in
  `docs/qa/reading-time-test-plan.md`).
- Attempted to run the Go toolchain in this sandbox (`go test ./...`);
  the command was blocked by this environment's own approval policy
  before it could execute. **No command actually ran, and no output was
  produced or seen.** This matches what earlier stages of this same flow
  also reported: the toolchain is not runnable on this box.

## Results

| # | Acceptance criterion | Covered by | Result | Evidence |
|---|------------------------|------------|--------|----------|
| 1 | ReadingTime helper exists in the text package | `text/reading_time.go` | ⬜ to be confirmed by a person | — |
| 2 | Blank/whitespace-only text → 0 minutes | test rows 1–2 | ⬜ to be confirmed by a person | — |
| 3 | Any other non-empty text → at least 1 minute | test rows 3–4 | ⬜ to be confirmed by a person | — |
| 4 | Longer text → more minutes, ceiling-rounded, 200 wpm | test rows 5–7 | ⬜ to be confirmed by a person | — |
| 5 | Speed fixed, no CJK special-casing | code read (no rate parameter or per-script branch found) | ⬜ to be confirmed by a person | — |
| 6 | Tests cover all four categories, full suite passes | `text/reading_time_test.go` + `go vet ./... && go test ./...` | ⬜ to be confirmed by a person — **not run on this box** | — |
| 7 | README lists the new files | `README.md` row | ⬜ to be confirmed by a person | — |

Legend: ⬜ to be confirmed by a person · ✅ verified (with evidence) ·
❌ failed (with evidence) · ⚠️ not verified (with reason)

## What a person needs to do to complete this report

1. Check out the commit on this flow branch that carries
   `text/reading_time.go`, `text/reading_time_test.go`, and the
   `README.md` update.
2. Run `go vet ./... && go test ./...` from the repository root (or read
   the CI run already attached to the relevant PR/commit) and record the
   literal pass/fail output.
3. For rows 1–4, 6, and 7 above, change ⬜ to ✅ or ❌ based on that run
   or a direct read of the named file, and paste the command output (or
   a link to the CI run, plus its result line) as evidence.
4. For row 5 (fixed speed, no CJK special-casing), confirm by reading
   `text/reading_time.go` that `ReadingTime` takes only `s string` and
   that the 200-words-per-minute rate is not exposed as a parameter or
   config value — this is a code-review confirmation, not something
   `go test` asserts on its own, so record it as a manual read rather
   than a test result.
5. If anything turns out ❌, this pass does not file a defect or open
   any issue/PR itself — record the failure here, in this file, for the
   next step of the flow to act on.

## Sprint verification queue note

`docs/sprint/reading-time.md` already recorded an empty verification
queue for this sprint. This report does not change that: it exists so a
person has one concrete checklist to tick off against a real
`go vet`/`go test` run (or the CI result), turning "CI will prove it"
into something actually confirmed rather than assumed.
