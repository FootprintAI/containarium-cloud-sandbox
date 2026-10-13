# PRD: Slugify support for non-ASCII letters

**Date:** 2026-10-14
**Status:** draft
**Owner:** unassigned

## Problem

`text.Slugify` turns every character that is not an ASCII letter or digit
into a separator. For titles written only in ASCII this is fine, but for
titles containing accented Latin letters or non-Latin scripts it destroys
the content of the slug instead of just the punctuation:

- `"Café déjà vu"` → every accented letter (`é`, `à`) is replaced by `-`,
  so the slug keeps only `caf`, `d`, `j`, `vu` separated by dashes — the
  words are mangled and barely recognizable.
- `"日本語のタイトル"` → **every** character is non-ASCII, so the slug is
  either empty or a single dash. Two completely different Japanese titles
  would collide on the same (empty) slug, which breaks any use of the slug
  as a unique, human-meaningful identifier (e.g., a URL path segment).

Who hurts: any engineering team using `Slugify` to build URL paths, file
names, or IDs from user- or author-supplied titles in a product with an
international audience. The cost is broken or colliding URLs and slugs
that give a reader no clue what the content is.

**Evidence:** the issue gives two concrete example titles but no data on
how often non-ASCII titles actually occur in production, or which scripts
dominate. This is an assumption to validate (see Open questions), not
measured evidence.

## Target user

A developer on a team that generates URL slugs or similar identifiers from
free-text titles, where some of those titles are not pure ASCII (product
names, article titles, user-submitted content, etc.).

## Success metrics

| Metric | Baseline | Target |
|--------|----------|--------|
| % of `Slugify` outputs that are empty or a single `-` for inputs containing at least one non-ASCII letter | unknown — instrument first | 0% for Latin-script accented input (MVP); tracked, not yet 0%, for non-Latin scripts (later phase) |
| Collision rate between two distinct titles that differ only in accented/non-Latin letters | unknown — instrument first | no regression vs. current behavior; ideally trending down as later phases land |

## MVP scope — the core journey

**Core journey:** a title containing accented Latin letters produces a
non-empty, readable, all-ASCII slug that a reader can still recognize as
related to the original title.

**Story:** As a developer generating a slug from a user- or
author-supplied title, I want accented Latin letters transliterated to
their closest plain-ASCII letter (e.g. `é` → `e`, `ñ` → `n`, `ü` → `u`) so
that the slug stays non-empty, readable, and URL-safe instead of being
chopped into dashes.
**Acceptance criteria:**
- [ ] `Slugify("Café déjà vu")` returns a non-empty, all-ASCII slug that
      reads as `cafe-deja-vu`
- [ ] A letter with no common single-letter ASCII equivalent keeps today's
      behavior (treated as a separator, collapsed per issue #3)
- [ ] All existing ASCII-only cases in `text/slug_test.go` are unaffected
**Priority:** P0 (MVP)

This MVP deliberately covers only Latin-script diacritics (the first
example in the issue), because that case has a well-known, mostly 1:1
mapping to ASCII (strip the combining accent) and is the smallest change
that stops the most common failure mode (mangled, not empty, slugs).

## Later phases

- **Non-Latin script support** (the `"日本語のタイトル"` case): Chinese,
  Japanese, Korean, Arabic, Cyrillic, etc. have no 1:1 letter-for-letter
  ASCII mapping — handling them means picking a romanization scheme per
  script (e.g., Hepburn for Japanese, Pinyin for Chinese) or deciding to
  keep/percent-encode native characters instead. This is materially more
  work than Latin diacritics and needs its own design pass
  (`/architect-design`) once there's a decision on direction (see Open
  questions). P1.
- **Configurable slug policy** (keep vs. transliterate vs. drop, chosen by
  the caller) if different consumers of `Slugify` turn out to need
  different behavior. Not pursued until a second consumer with different
  needs actually exists. P2.

## Out of scope

- **Silently dropping non-ASCII letters (today's behavior).** Explicitly
  rejected as the long-term answer: for fully non-Latin titles it produces
  an empty or near-empty slug, which breaks uniqueness and readability —
  the exact problem this PRD exists to fix.
- **Keeping raw Unicode letters in the slug.** Rejected for the MVP: the
  function's own contract (see its doc comment) is to produce a
  "URL-safe slug," which this repo's codebase and tests assume is ASCII;
  switching to raw Unicode output would be a breaking change to that
  contract and needs its own product decision, not a side effect of this
  fix.

## Open questions & assumptions

- **Assumption (unvalidated):** accented Latin-script titles are the more
  common real-world case motivating this issue; the non-Latin example is
  evidence that the need exists but not of its volume. Action: instrument
  `Slugify` call sites (or their inputs) to measure how often non-ASCII
  titles occur and which scripts appear, before investing in the later
  phases.
- **Open question (needs a person's decision, not assumed here):** for
  non-Latin scripts, should the long-term answer be romanization,
  preserving native characters (e.g., percent-encoded in the slug), or
  something else? This PRD does not decide it — it is explicitly deferred
  to a later phase pending real usage evidence.
- **Open question (engineering, not product):** which transliteration
  table/library should back the Latin-diacritic MVP (e.g., a small
  hand-written mapping vs. Unicode normalization + combining-mark
  stripping)? Left to `/architect-design` once this PRD is accepted.

## Hand-off note

Under normal operation this PRD would now hand off to delivery: one
GitHub issue per P0 story, labeled `product`, with the acceptance criteria
above in the body, so `/pm-sprint-delivery` can schedule it. That issue
was not filed because this run's platform rules route output into this
document instead of creating issues/PRs/comments. The P0 story and its
acceptance criteria above are what would have been filed.

> PRD written to `docs/product/slugify-non-ascii-support.md`: 1 P0 story
> (Latin-diacritic transliteration), 2 deferred (non-Latin scripts,
> configurable policy).
> Success metric: % of degenerate (empty/single-dash) slugs for non-ASCII
> input, from unknown baseline to 0% for Latin-script accented input.
> Next: `/architect-design slugify-non-ascii-support` for the
> transliteration-table design, then `/pm-sprint-delivery` to schedule the
> P0 story.
