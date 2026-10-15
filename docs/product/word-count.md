# PRD: WordCount helper for the text package

**Date:** 2026-10-15
**Status:** draft
**Owner:** unassigned

## Problem

There is no shared way in this codebase to count the words in a piece of
text. A developer building a feature that needs a word count — a
content-length label, a minimum/maximum word-count validation on a form
field, a metric derived from text length — has to write that
word-splitting logic themselves, with no guarantee two call sites agree
on what counts as a "word."

Who hurts: a developer who needs a word count for a piece of text. The
cost today is either skipping the feature or duplicating slightly
different word-splitting logic per call site.

**Evidence:** the issue itself labels this "a fixture for the flows
opt-in test" — it exists to exercise this repository's agent pipeline,
not because a real reported problem or usage data motivated it. The
issue does already give a complete, unambiguous definition (`WordCount`
is the number of runs of non-space characters), so there is no
functional ambiguity to resolve — just a feature request to scope.
Treated here as a straightforward, already-specified request rather than
something needing more problem validation.

## Target user

A developer who needs to count the words in a piece of text and wants a
single, shared definition of "word" instead of writing their own
splitting logic.

## Success metrics

Because this is a small, internal utility function rather than a
user-facing feature with its own usage funnel, these are implementation
and consistency metrics rather than engagement ones:

| Metric | Baseline | Target |
|--------|----------|--------|
| % of `WordCount` outputs matching the spec'd definition (count of maximal runs of non-space characters) across a representative test corpus | unknown — no implementation exists yet | 100%, proven by the unit tests written in the implementation task |
| Number of call sites computing a word count with their own ad hoc logic instead of this helper | unknown — not currently tracked | 0 for any new call site added after this ships (existing duplicates, if any are found, are a separate cleanup, not part of this MVP) |

## MVP scope — the core journey

**Core journey:** given any string, get back the exact count of its
maximal runs of non-space characters.

**Story:** As a developer who needs to count the words in a piece of
text, I want `text.WordCount(s string) int` to count runs of non-space
characters, so that I don't have to reimplement word-splitting logic
myself.
**Acceptance criteria:**
- [ ] `WordCount("")` returns `0`
- [ ] A whitespace-only string (spaces, tabs, newlines) returns `0`
- [ ] A single-word string returns `1`
- [ ] A string of N words separated by any run of whitespace (single
      spaces, multiple spaces, tabs, newlines) returns `N`
- [ ] A run of non-space characters that includes punctuation (e.g.
      `"word,"`) counts as one word, exactly as the definition implies —
      this is the spec working as intended, not a gap
**Priority:** P0 (MVP)

This MVP is already the complete feature: the issue's own definition
("the number of runs of non-space characters") leaves nothing
underspecified for a space-delimited count.

## Later phases

- **Shared word-counting logic.** If another helper in this package ever
  needs its own internal word count, it should call `WordCount` rather
  than re-deriving the same split — but that is a refactor question for
  whichever change introduces the second use, not a reason to delay or
  complicate this one. Not scoped here; left for `/architect-design` to
  raise if and when it applies.

## Out of scope

- **Script-aware or semantic word counting** (e.g., treating a run of
  CJK characters with no internal spaces as more than one "word"). The
  issue defines `WordCount` exactly as a whitespace-run counter;
  anything smarter is a different, unrequested feature, not a cut
  corner of this one.
- **Stripping punctuation before counting.** Not requested, and doing so
  would contradict the issue's literal definition (a run of non-space
  characters, punctuation included, is one word as specified).

## Open questions & assumptions

None. The issue's definition — "the number of runs of non-space
characters" — is complete and unambiguous on its own: it fixes what a
word is, requires no reading-speed-style rate or configuration, and
behaves the same for every script without a separate convention to
invent (unlike word-counting features that estimate something, such as
reading time, this one counts something, so there is no analogous
"CJK" or "configurable rate" question to carry forward). This is a
legitimate "no open questions" result, not an unexamined scope.

## Hand-off note

Under normal operation this PRD would now hand off to delivery: one
GitHub issue per P0 story, labeled `product`, with the acceptance
criteria above in the body, so `/pm-sprint-delivery` can pull it into an
umbrella sprint issue. No such issue was filed because this run's
platform rules route output into this document instead of creating
issues/PRs/comments.

> PRD written to `docs/product/word-count.md`: 1 P0 story (whitespace-run
> word counting), 0 deferred (the MVP is the complete feature; see Later
> phases for a non-blocking refactor note).
> Success metric: WordCount output matches the spec'd definition on 100%
> of the test corpus (no baseline exists yet — nothing has been built).
> Next: `/architect-design word-count` for the technical design, then
> `/pm-sprint-delivery` to schedule the P0 story.
