# Sprint plan: ship the ReadingTime helper

**Date:** 2026-10-15
**Status:** proposed
**Owner:** unassigned
**Flow:** f16c5d49, entering directly at the sprint stage — the issue's
own "Agreed scope" section fixes the behavior, so no separate product or
architecture stage precedes this one.

## Goal

By the end of this sprint, `text.ReadingTime(s string) time.Duration`
exists in the `text` package exactly as scoped by the issue, is covered
by table-driven tests, and is mentioned in the repo's README.

## Scope, as agreed in the issue

- Estimate reading time at a fixed 200 words per minute, rounded up to a
  whole minute.
- At least one minute for any non-empty text.
- Zero for empty or whitespace-only text.
- A "word" is a run of non-space characters — this single rule is also
  what keeps non-Latin scripts out of the problem: there is no separate
  CJK convention to invent, because the counting rule doesn't change per
  script.
- Explicitly out of scope: a configurable reading speed, and
  character-based counting for CJK text.

## Phase 1 — make the flow work (this week)

- [ ] **Task 1 — Implement `text.ReadingTime`** — unassigned, not
  blocked
  - **Done means:** the function exists in the `text` package, matches
    the agreed scope above, and the worked examples hold: a 200-word
    input returns exactly 1 minute, a 400-word input returns exactly 2
    minutes, a 450-word input returns 3 minutes (the remainder rounds
    up).
  - **CI can prove this done:** yes — pure logic, no running environment
    needed.
- [ ] **Task 2 — Table tests for `ReadingTime`, plus a README line** —
  unassigned, blocked by Task 1 (the tests exercise the function Task 1
  adds)
  - **Done means:** table-driven tests cover the empty, whitespace-only,
    short-text, and the three worked-example cases above; the root
    `README.md`'s existing "What is in the repo" table gets one new row
    for the new files, in the same style as the existing row for
    `text/slug.go` / `text/slug_test.go`.
  - **CI can prove this done:** yes — pure logic tests, no running
    environment needed.

This is a two-task sprint by design, matching the issue's own expected
shape (one task for the function, one for its tests and the README
line) — there is nothing else on the critical path of this flow.

## Phase 2 — optimize (next phase, only after this flow works)

None proposed. There is no further optimization step for a single pure
helper function to schedule behind these two tasks this week.

## Out of sprint (deferred)

- **Configurable reading speed** and **character-based CJK counting** —
  both explicitly out of scope per the issue's own "Agreed scope"
  section. Not scheduled here because they were deliberately excluded,
  not because they were cut for time.

## Verification queue — batched, one deployed version at a time

**None.** Every acceptance criterion for both tasks is pure logic,
already named as a test case above — a green `go vet ./... && go test
./...` run proves them. No deployed environment is needed to confirm
this flow works.

**Deployed to dev:** n/a — nothing in this sprint needs a deployed
environment to prove done.
**Deployed to prod:** n/a

## Open questions — answered before the cycle runs unattended

None. Reading both tasks as their implementer would: the reading speed
is fixed by the issue (no configurability to design), the word-counting
rule is given directly ("runs of non-space characters," equivalent to
splitting on whitespace — this also resolves what would otherwise be a
CJK-handling question, since the same whitespace rule applies regardless
of script), the one-minute floor and zero cases are both spelled out,
and the task split (function / tests+README) is handed down by the issue
itself. No point was found where an implementer would have to guess.
This is a legitimate "no open questions" result, not an unexamined scope.

## Cadence

- Progress would normally be posted as comments on the umbrella issue
  for this sprint, once the platform files the two tasks declared with
  this plan.
- Definition of done for the sprint: `text.ReadingTime` exists, its test
  table passes, `go vet ./... && go test ./...` is green, and the README
  mentions the new files.

## Status

🟢 on track — scope is fully specified by the issue, both tasks are
small and CI-provable, and no open question blocks either one.
