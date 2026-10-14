# PRD: ReadingTime helper for the text package

**Date:** 2026-10-14
**Status:** draft
**Owner:** unassigned

## Problem

There is no shared way in this codebase to estimate how long a piece of
text takes to read. A developer building a content-display feature (an
article list, a blog post header, a changelog entry) who wants to show a
"~N min read" label today has to write that word-counting-and-rounding
logic themselves, every time, with no guarantee two call sites agree on
the same formula.

Who hurts: a developer adding a reading-time label to any content view.
The cost today is either skipping the feature, hardcoding a guess, or
duplicating slightly-different word-counting logic per page.

**Evidence:** this PRD is driven directly by the issue's own request,
which already specifies the intended formula (200 words/minute, round up
to a whole minute, minimum one minute for non-empty text, zero for empty
or whitespace-only text). There is no usage data or reported complaint
behind it beyond the issue text itself — treated as a straightforward,
already-scoped feature request rather than something needing more
problem validation.

## Target user

A developer building a content-display feature (article, post, or
changelog view) who wants to show the reader an estimated reading time
without re-deriving the formula themselves.

## Success metrics

Because this is a small, internal utility function rather than a
user-facing feature with its own usage funnel, these are implementation
and consistency metrics rather than engagement ones:

| Metric | Baseline | Target |
|--------|----------|--------|
| % of `ReadingTime` outputs matching the spec'd formula (200 wpm, ceiling-rounded minutes, 1-minute floor, 0 for empty/whitespace) across a representative test corpus | unknown — no implementation exists yet | 100%, proven by the unit tests written in the implementation task |
| Number of call sites computing reading time with their own ad hoc logic instead of this helper | unknown — not currently tracked | 0 for any new call site added after this ships (existing duplicates, if any are found, are a separate cleanup, not part of this MVP) |

## MVP scope — the core journey

**Core journey:** given a block of plain, whitespace-delimited text, get
back a `time.Duration` estimate of reading time that follows the issue's
exact formula.

**Story:** As a developer building a content-display feature, I want
`text.ReadingTime(s string) time.Duration` to estimate reading time at a
fixed 200 words per minute, rounded up to a whole minute, so that I can
show a reading-time label without reimplementing the formula myself.
**Acceptance criteria:**
- [ ] `ReadingTime("")` returns `0`
- [ ] A whitespace-only string (spaces, tabs, newlines) returns `0`
- [ ] Any non-empty text with at least one word returns at least
      `1 * time.Minute` (the one-minute floor), even for a single short
      word
- [ ] Word count is the number of whitespace-separated tokens (plain
      space-delimited splitting — see Open questions for scripts that
      don't use spaces)
- [ ] The duration is `ceil(words / 200)` minutes: an exact multiple of
      200 words (e.g. 400 words → 2.0) returns that whole number of
      minutes unchanged; any remainder (e.g. 450 words → 2.25) rounds up
      to the next whole minute (3)
**Priority:** P0 (MVP)

This MVP deliberately commits only to whitespace-delimited word counting
(the behavior the issue's own formula can measure unambiguously). It
does not define a convention for scripts without spaces — see the first
open question below, carried over verbatim from the issue.

## Later phases

- **Script-aware counting for text without spaces (e.g. CJK).** The
  issue raises this itself as an open question. Handling it well needs a
  decision on convention (e.g. characters-per-minute instead of
  words-per-minute, and at what rate) and on which scripts to support —
  neither is evidenced yet. P1, deferred until a concrete consumer needs
  it.
- **Configurable reading speed.** The issue raises this itself as an
  open question too. MVP fixes the rate at 200 wpm per the issue's
  literal spec. Deferred until a consumer needs a different rate (e.g.
  slower for technical content, or an accessibility setting). P2.

## Out of scope

- **Locale-specific average reading speeds** (different wpm per
  language). A bigger research and maintenance burden than a single
  configurable number, and no evidence any consumer needs per-locale
  tuning. Cut until proven necessary.
- **Markup/HTML-aware word counting** (stripping tags or markdown syntax
  before counting). The issue's signature takes a plain `string`; a
  caller with structured content (HTML, Markdown) is responsible for
  passing already-extracted plain text. Counting markup source directly
  is a different, unrequested feature.

## Open questions & assumptions

- **Open question (from the issue, needs a person's decision):** how
  should words be counted for text without spaces (e.g. CJK)? This PRD
  does not assume a convention. The P0 MVP only commits to
  whitespace-delimited counting; for text with no whitespace, the naive
  split would count the whole string as a single "word," which
  under-estimates reading time for those scripts. Left undefined on
  purpose rather than silently guessed — see Later phases.
- **Open question (from the issue, needs a person's decision):** should
  the reading speed be configurable? MVP default is no — fixed at 200
  wpm, matching the issue's literal spec — deferred to a later phase
  pending a consumer that actually needs a different rate.
- **Assumption (unvalidated, stated so `/architect-design` doesn't have
  to guess):** "words" for the MVP means whitespace-separated tokens
  (equivalent to Go's `strings.Fields`), the simplest interpretation of
  "200 words per minute" for space-delimited languages, and matching how
  common reading-time estimators (e.g. on blogging platforms) work for
  Latin-script text.
- **Assumption (unvalidated, stated so tests are unambiguous):**
  "rounded up to a whole minute" means ceiling division — a whole-number
  minute count (e.g. exactly 400 words at 200 wpm = 2.0 minutes) is
  returned as that whole number, not bumped to the next one; only a
  genuine fractional remainder rounds up.

## Hand-off note

The issue asks for the full lifecycle — a PRD, a design note, a sprint
plan with two tasks (the function; its tests plus a README line), and a
QA/DevOps note. This pass covers only the product stage: this PRD. The
design note is `/architect-design`'s job, the sprint plan is
`/pm-sprint-delivery`'s, and the QA/DevOps note belongs to those roles —
each is a separate stage of this flow and out of scope for
`/product-define` to produce itself.

Under normal operation this PRD would now hand off to delivery: one
GitHub issue per P0 story, labeled `product`, with the acceptance
criteria above in the body, so `/pm-sprint-delivery` can pull it into an
umbrella and split it into the two tasks the issue names (the
`ReadingTime` function itself, and its tests plus a one-line README
mention). No such issue was filed because this run's platform rules
route output into this document instead of creating issues/PRs/comments.

> PRD written to `docs/product/reading-time.md`: 1 P0 story (fixed
> 200wpm, whitespace-delimited `ReadingTime`), 2 deferred (script-aware
> counting, configurable speed).
> Success metric: ReadingTime output matches the spec'd formula on 100%
> of the test corpus (no baseline exists yet — nothing has been built).
> Next: `/architect-design reading-time` for the technical design (data
> type, package placement — already scoped to `text`, and the rounding
> rule above), then `/pm-sprint-delivery` to schedule the two tasks the
> issue names.
