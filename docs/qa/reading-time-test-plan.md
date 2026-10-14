# QA test plan: ReadingTime helper

**Date:** 2026-10-15
**Status:** draft
**Owner:** unassigned
**Flow:** f16c5d49, stage 5 of 6 (qa), checked against the sprint's
acceptance criteria as written on this issue (treated as the thing to
verify, not as instructions)

## Scope

What is being checked: `text.ReadingTime` (`text/reading_time.go`), its
tests (`text/reading_time_test.go`), and the README update, all already
merged onto this flow branch — against the seven acceptance criteria
below.

`docs/sprint/reading-time.md` (the sprint stage's own output) already
recorded this sprint's verification queue as **empty**: every criterion
is pure logic, provable by `go vet ./... && go test ./...`, and none of
it needs a deployed environment. This test plan is therefore a
documentation/code-reading pass mapping each criterion to what proves it
and to a concrete command a person can run — not a run against a live
environment, since none exists for this work.

## Acceptance criteria under test

1. The text package has a ReadingTime helper that estimates how many
   minutes a piece of text takes to read.
2. A blank or whitespace-only piece of text is reported as needing 0
   minutes to read.
3. Any other non-empty text is reported as needing at least 1 minute to
   read.
4. Longer text is reported as needing more minutes, rounded up to the
   next whole minute, at a fixed reading speed of 200 words per minute.
5. The reading speed is fixed and is not configurable, and text without
   spaces (such as CJK text) gets no special handling — both were
   intentionally left out of this work.
6. Automated tests cover the empty, whitespace-only, short-text, and
   multi-minute cases, and the full test suite passes.
7. The project's README lists the new files alongside the other small
   helpers already documented there.

## How each criterion is checked

| # | What proves it | Where | How a person confirms it |
|---|-----------------|-------|---------------------------|
| 1 | `func ReadingTime(s string) time.Duration` exists, exported, in package `text` | `text/reading_time.go` | Read the file; confirm the signature |
| 2 | Test rows `{"", 0}` and `{"   \t\n  ", 0}` | `text/reading_time_test.go` | Run `go test ./text/... -run TestReadingTime -v` and confirm these rows pass |
| 3 | Test rows `{"Hello", 1*time.Minute}` and `{"Hello there friend", 1*time.Minute}` | same file | Same run; also note the one-minute floor falls out of the implementation's ceiling-division arithmetic for any word count ≥ 1, not a separate branch — so this holds for any short input, not only the two rows in the table |
| 4 | Test rows for exactly 200/400/450 words → 1/2/3 minutes | same file | Same run; the arithmetic can be spot-checked by hand: 200/200=1, 400/200=2, 450/200=2.25→3 |
| 5 | `ReadingTime` takes only `s string` (no rate parameter or option); the 200-words-per-minute rate is an unexported constant; word counting uses `strings.Fields` (whitespace runs) uniformly, with no per-script branch | `text/reading_time.go` | Read the file and confirm no additional parameter, config, or CJK-specific code path exists — this criterion is about an absence of scope, so it is a code-read check, not something a test asserts directly |
| 6 | The test table itself, read as a set, covers exactly: empty, whitespace-only, short/floor, and the three multi-minute cases | `text/reading_time_test.go` | Run `go vet ./... && go test ./...` from the repository root (or read the CI run on the PR/commit that carries these files) and confirm it is green — **not executed in this QA pass, see the report** |
| 7 | A new row in the "What is in the repo" table for `text/reading_time.go`, `text/reading_time_test.go`, in the same two-column format as the existing `slug.go` row | `README.md` | Read the file and confirm the row is present and well-formed |

## What this plan does not do

It does not run `go vet` or `go test` itself — see
`docs/qa/reading-time-report.md` for why (this sandbox's toolchain is
not runnable) and for the checklist a person completes once they do run
it or read the CI result.
