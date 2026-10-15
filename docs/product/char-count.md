# PRD: Character count helper for the `text` package

**Date:** 2026-10-15
**Status:** draft
**Owner:** unassigned

## Problem

Issue #50 asks for a way to count the characters in a string correctly when
it contains multi-byte characters (for example, an accented letter). Go's
native `len(s)` counts bytes, not runes, so any caller that measures a
string's visible length today either re-implements rune counting ad hoc or
gets a wrong count for non-ASCII input.

The issue itself is the only evidence available — no linked ticket, caller,
or user complaint is cited. That is stated here as an assumption, not
dressed up as research: the need is for a small, correct, rune-counting
utility consistent with how `text.Slugify` already walks a string rune by
rune, not a response to a measured user complaint.

## Target user

Developers in this repository who import the `text` package and need the
number of characters (not bytes) in a string — for example, to validate
input length or truncate a string for display.

## Success metrics

This is a small internal helper in a library package, not an end-user
feature, so activation/completion-style metrics don't apply. The bar is
correctness:

| Metric | Baseline | Target |
|--------|----------|--------|
| Correct rune count on multi-byte input | 0 (no such helper exists today) | `CharCount` returns the exact rune count for ASCII input, multi-byte input (e.g. an accented letter), and the empty string, proven by table tests |

## MVP scope — the core journey

The one journey: a developer calls `text.CharCount(s)` and gets back the
number of characters (runes) in `s`, so a string with one accented letter
counts as one character rather than two bytes.

**Story 1:** As a developer using the `text` package, I want a
`CharCount(s string) int` function so that I can get the number of
characters in a string instead of its byte length.
**Acceptance criteria:**
- [ ] `CharCount` counts runes, not bytes — a string containing one
  multi-byte character (e.g. an accented letter such as `"é"`) returns `1`,
  not `2`
- [ ] `CharCount("")` returns `0`
- [ ] Plain ASCII input returns the same value `len(s)` would give (e.g.
  `CharCount("abc")` returns `3`)
**Priority:** P0 (MVP)

**Story 2:** As a maintainer of the `text` package, I want table-driven
tests for `CharCount` and a one-line README mention so that its rune-vs-byte
behaviour is documented and protected against regression.
**Acceptance criteria:**
- [ ] A table test covers at least: the empty string, an ASCII-only string,
  and a string containing a multi-byte character
- [ ] `go test ./...` passes
- [ ] The README's file/purpose table gains one line describing the new
  file(s)
**Priority:** P0 (MVP)

## Later phases

None. The issue scopes this to one function plus its tests and a doc line;
there is no fast-follow capability requested.

## Out of scope

- **Grapheme-cluster-aware counting** (treating a combining-mark sequence or
  a modifier emoji as one visual character the way a reader would) — the
  issue asks for a rune count, which is what Go's `range` over a string
  already yields; grapheme segmentation needs a dedicated library and no
  evidence in the issue calls for it.
- **Changes to `text.Slugify` or any other existing function** — unrelated
  to this request.

## Open questions & assumptions

- **Assumption:** "characters" in the issue means Unicode code points
  (runes) — what `CharCount` would return by ranging over the string in Go —
  not grapheme clusters. The issue's own example (a precomposed accented
  letter counting as one) is consistent with this reading but doesn't
  exercise a combining-mark sequence, so the reading is an assumption, not
  something the issue confirms.
- **No additional evidence** (ticket, caller, support complaint) exists
  beyond the issue text itself.

## Hand-off

This step normally opens one GitHub issue per P0 story, labeled `product`,
for sprint planning. For this change that step is recorded here instead:
the two P0 stories above are ready to hand to an engineering implementation
pass as the two small tasks the issue names:

1. Add `CharCount(s string) int` to the `text` package.
2. Add its table-driven tests and the one-line README mention.
