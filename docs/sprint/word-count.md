# Sprint plan: ship the WordCount helper

**Date:** 2026-10-15
**Status:** proposed
**Owner:** unassigned
**Flow:** 87545650, stage 3 of 6 (sprint), building on
`docs/product/word-count.md` and `docs/architecture/word-count.md`

## Goal

By the end of this sprint, `text.WordCount(s string) int` exists exactly
as specified in the design doc, is covered by its full test table, and
is mentioned in the repo's README — the whole feature from the product
PRD works end-to-end.

## Phase 1 — make the flow work (this week)

- [ ] **Task A — Implement `text.WordCount`** — unassigned, not blocked
  - **Done means:** `text/word_count.go` exists with the exact function
    from `docs/architecture/word-count.md` (`len(strings.Fields(s))`) —
    `WordCount("")` and whitespace-only input return `0`; a single-word
    input returns `1`; N words separated by any run of whitespace
    (spaces, tabs, newlines) return `N`; a punctuation-attached word
    (e.g. `"word,"`) counts as one word
  - **CI can prove this done:** yes — pure logic, no running environment
    needed
- [ ] **Task B — Tests for `WordCount`, plus a README line** —
  unassigned, blocked by Task A (the tests exercise the function Task A
  adds)
  - **Done means:** `text/word_count_test.go` added with the full table
    from the design doc's Test strategy section (empty, whitespace-only,
    single word, multiple words with single spaces, multiple words with
    mixed whitespace, punctuation-attached words); `go vet ./... && go
    test ./...` stays green repo-wide; the root `README.md`'s existing
    "What is in the repo" table gets one new row for
    `text/word_count.go`, `text/word_count_test.go`, in the same
    two-column format as the existing `slug.go` row
  - **CI can prove this done:** yes — pure logic tests, no running
    environment needed

Both tasks are on the critical path of this flow's only deliverable;
there is nothing else to schedule within Phase 1 — this is a two-task
sprint by design, matching the established pattern for a small `text`
package helper in this repo.

## Task list (as filed with the platform)

This is the literal title and body of each task declared alongside this
plan — the platform files these, not a summary of them. If this section
and the "tasks" list diverge, this section is stale and the filed tasks
are the source of truth.

### Task 1 — Implement text.WordCount in the text package

> Add WordCount(s string) int to the text package. Return the number of
> maximal runs of non-space characters in s -- implement this as a
> single call to strings.Fields and return the length of its result,
> since that function's output already is exactly this count (it splits
> on any run of whitespace and discards empty results from
> leading/trailing whitespace). An empty string or a string containing
> only whitespace returns 0. A single word returns 1. A run of
> non-space characters that includes punctuation (for example "word,")
> counts as one word, exactly as the definition implies. Verified by:
> the full test suite passing (go vet ./... && go test ./... stays
> green across the repo), and the behavior matching the worked cases --
> an empty or whitespace-only string returns 0, a single word returns
> 1, and a string of several words separated by any mix of spaces,
> tabs, or blank lines returns the exact number of words.

### Task 2 — Add table tests for WordCount and a README line

> Add table-driven tests for text.WordCount in the text package,
> covering: an empty string returns 0; a whitespace-only string (spaces,
> tabs, newlines) returns 0; a single word returns 1; a string of
> several words separated by single spaces returns the matching count;
> a string of several words separated by a mix of whitespace (tabs and
> blank lines as well as spaces) returns the matching count; and a
> string with punctuation attached to words (for example "word, word.")
> counts each as one word. This task depends on the WordCount function
> already existing in the same package from the other task. Also add
> one row to the root README.md's existing "What is in the repo" table
> describing the new files, following the same style as the existing
> row for text/slug.go and text/slug_test.go (a short one-sentence
> description of what the files contain). Verified by: the full test
> suite passing (go vet ./... && go test ./... stays green across the
> repo) with the new tests included, and the new README row present and
> following the existing table's two-column format.

## Phase 2 — optimize (next phase, only after this flow works)

None proposed. The flow is "the function exists, is tested, and is
documented" — there is no further optimization step for a single pure
helper function to schedule behind it this week.

## Out of sprint (deferred)

- **A shared word-splitting helper reused by `WordCount` and any other
  helper that needs a word count** — already deferred by both the PRD
  and the design doc (`docs/product/word-count.md` and
  `docs/architecture/word-count.md`, both under "Later phases" /
  "Rejected alternatives") pending a second consumer actually needing
  one. Not scheduled here because it isn't needed yet, not because it
  was cut for time.

## Verification queue — batched, one deployed version at a time

**None.** Every acceptance criterion for both tasks is pure-logic and
already covered by the unit test table in the design doc — a green
`go test ./...` run proves them. No `needs-verification` label applies
to either task, and no deployed environment is needed to confirm this
flow works.

**Deployed to dev:** n/a — nothing in this sprint needs a deployed
environment to prove done.
**Deployed to prod:** n/a

## Open questions — answered before the cycle runs unattended

None. Reading both tasks as their implementer would: the function body
is fully specified in the design doc (a single call to
`strings.Fields`), every behavior the PRD names (empty, whitespace-only,
single word, N words, punctuation-attached words) already has a named
test row with a concrete input/expected pair, the file names and README
row format are fixed by convention (mirroring `slug.go`), and CI already
runs `go vet ./... && go test ./...` on the whole module (confirmed in
`.github/workflows/ci.yml`), so no CI configuration change is needed for
either task. No point was found where an implementer would have to
guess. This is a legitimate "no open questions" result, not an
unexamined scope.

## Cadence

- Progress would normally be posted as comments on the umbrella issue,
  weekly or on status change. No GitHub issue was created for this
  sprint — see Hand-off note.
- Definition of done for the sprint: `text.WordCount` exists, its full
  test table passes, `go vet ./... && go test ./...` is green, and the
  README mentions it.

## Status

🟢 on track — scope is fully specified by the PRD and design doc, both
tasks are small and CI-provable, and no open question blocks either one.

## Hand-off note

Under normal operation this stage would open one GitHub umbrella issue
(the content above) plus one child issue per task (Task A, Task B),
labeled and cross-linked back to the umbrella. None of that was created
because this run's platform rules route output into this document, and
into the structured task list declared alongside it (reproduced verbatim
in "Task list (as filed with the platform)" above for review), instead
of creating issues/PRs/comments.

> Sprint created: 2 Phase 1 tasks (flow), 0 Phase 2 tasks. Critical path:
> Task A → Task B.
> Verification queue: empty — both tasks are CI-provable.
> Open questions: none found.
> Next: `/engineer-implement` picks up Task A, then Task B.
