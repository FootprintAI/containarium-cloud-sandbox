# Sprint 2026-W42: Character count helper for the `text` package

**Date:** 2026-10-15
**Status:** 🟢 on track
**Owner:** unassigned

## Goal

By the end of the week, a developer can call `text.CharCount(s)` and get the
correct number of characters (Unicode code points) in `s`, including strings
with multi-byte characters, backed by a green test suite.

This sprint plans the work scoped by the earlier stages of this flow:
`docs/product/char-count.md` (PRD, issue #50) and
`docs/architecture/char-count.md` (design). Both are already fully decided —
there is nothing left to re-litigate here, only to sequence into
independently implementable tasks.

## Phase 1 — make the flow work (this week)

The whole sprint is Phase 1: the flow is a single small function plus its
tests, and it either works correctly or it doesn't — there is no smaller
"ugly but working" version and no optimization pass that would follow it.

- [ ] **Task 1** — Add `text.CharCount(s string) int`
- [ ] **Task 2** — Add table-driven tests for `CharCount` and the README
  entry (depends on Task 1 existing to test against)

## Phase 2 — optimize (next phase, only after flow works)

None. The design doc already rejected the only two alternative shapes
(hand-rolled rune-counting loop, grapheme-cluster-aware counting) as
unnecessary for this change, so there is nothing to defer into a later
optimization pass.

## Out of sprint (deferred)

- **Grapheme-cluster-aware counting** — explicitly out of scope per the PRD;
  not a planned follow-up, just not what was asked for.
- **Changes to `text.Slugify` or any other existing function** — unrelated
  to this request.

## Verification queue — batched, one deployed version at a time

Empty. Both tasks are pure logic in a library package with no network
boundary, no UI, no migration, and no third-party integration — every
acceptance criterion is provable by `go vet ./...` and `go test ./...`,
which are already wired in `.github/workflows/ci.yml`. Nothing here needs a
deployed environment or a batched verification pass.

## Open questions — answered before the cycle runs unattended

None found. Both the PRD and the design doc already resolved the one place
an implementer could have guessed — whether "characters" means Unicode code
points (runes) or grapheme clusters — in favor of runes, matching how
`text.Slugify` already walks a string, and documented it as a stated
assumption rather than leaving it open. Nothing else in scope has more than
one reasonable way to build it, and no dependency outside the repo is
involved.

## Cadence

- This is a one-week sprint; no weekly progress update is expected beyond
  the sprint's own close, since both tasks are small enough to land within
  the week.
- Definition of done: `text.CharCount` exists, is covered by table-driven
  tests for the empty string, ASCII input, and multi-byte input, the
  README's file/purpose table mentions the new files, and `go vet ./...`
  and `go test ./...` both pass.

## Status

🟢 on track — scope is small, fully specified by the prior stages, and
carries no open questions or verification items.
