# PRD: Cap the length of Slugify output

**Date:** 2026-10-14
**Status:** draft
**Owner:** unassigned

## Problem

`text.Slugify` has no maximum output length today. A long title produces
an equally long slug, with no limit. Long slugs are awkward wherever they
end up — harder to read in a URL, and some downstream systems (database
columns, path-length limits, link-sharing UIs) may reject or mangle a
slug past a certain length.

Who hurts: the same developer who already relies on `Slugify` to build a
URL path, file name, or ID from a title (see
`docs/product/slugify-non-ascii-support.md`, from issue #5) — specifically
when that title is unusually long.

**Evidence:** the issue states this as a hypothetical risk ("may hit
limits in some systems") rather than a reported incident or a cited
system constraint. No production data on actual slug lengths, and no
named downstream system with a known limit, is available. This PRD
treats the need for a cap as plausible but unvalidated, and says so
explicitly rather than inventing a number (see Open questions).

## Target user

Same as the non-ASCII PRD: a developer on a team that generates URL slugs
or similar identifiers from free-text titles — here, specifically one
whose titles can run long (e.g., full article headlines, descriptive
product names).

## Success metrics

| Metric | Baseline | Target |
|--------|----------|--------|
| Distribution of slug lengths (p50 / p95 / max) for representative real titles | unknown — instrument/sample first | gathered before a cap value is fixed, so the cap reflects real usage instead of a guess |
| % of generated slugs ending in a dangling `-` after a cut | not applicable today (no cap exists) | 0% once capping ships |

## MVP scope — the core journey

**Core journey:** a title long enough that its slug would otherwise be
unbounded instead produces a slug capped at a bounded maximum length,
cut at a separator boundary rather than mid-word where possible, with no
trailing dash left behind.

**Story:** As a developer generating a slug from a long title, I want
`Slugify`'s output capped at a bounded maximum length — cut at a `-`
boundary rather than mid-word when one is available, and with any
trailing `-` left by the cut removed — so the slug stays usable in
systems that have length limits.
**Acceptance criteria:**
- [ ] A title whose generated slug would otherwise exceed the cap
      produces a slug no longer than the cap
- [ ] The cut prefers the last `-` boundary at or before the cap, so a
      word isn't split, when such a boundary exists within the cap
- [ ] A trailing `-` left by the cut (from a boundary cut or a hard cut)
      is trimmed — mirroring the trim already applied to leading/trailing
      `-` by the issue #3 fix — so a capped slug never ends in `-`
- [ ] If no `-` boundary exists within the cap (a single very long word),
      the slug is hard-cut at the cap and any resulting trailing `-` is
      still trimmed
- [ ] A slug already at or under the cap is returned unchanged
- [ ] Existing cases in `text/slug_test.go` (all well under any
      reasonable cap) are unaffected
**Priority:** P0 (MVP) — implementation is blocked on Open questions #1
(the cap value) and #2 (the API shape) below being answered by a person;
this PRD defines the mechanism, not the number.

## Later phases

- **Script-aware / display-width length limits.** Once non-ASCII support
  (the P1 phase of `docs/product/slugify-non-ascii-support.md`) ships,
  a plain rune count may no longer match what "too long" means for every
  script. Not pursued until that phase lands and turns out to need it.
- **Per-consumer configurable cap**, if a second consumer of `Slugify`
  turns out to need a different limit than the first. Mirrors the
  "configurable slug policy" later-phase item already deferred in the
  non-ASCII PRD for the same reason: don't build configurability before a
  second need for it exists.

## Out of scope

- **Adding a truncation marker (e.g. `…` or `-more`) to a capped slug.**
  The issue asks only about the trailing dash left by the cut, not about
  signaling "this was shortened" inside the slug itself. A slug is an
  identifier, not display text — unlike `text.Truncate` (issue #7), which
  exists precisely to add a visible "shortened" marker for display
  titles. Mixing that concern into `Slugify` would make capped slugs
  longer than necessary for a benefit nobody asked for here.

## Open questions & assumptions

- **Open question (needs a person's decision — the central question of
  this issue):** what should the cap value be? The issue gives no cited
  system limit or production data. Common reference points exist (e.g.,
  SEO guidance often suggests keeping URL slugs under ~60 characters for
  readability; many database slug/identifier columns allow well beyond
  that), but picking one without knowing which downstream systems
  actually consume these slugs would be inventing evidence. Recommended
  next step: sample real titles (see Success metrics) and confirm any
  concrete downstream constraint before fixing a number.
- **Open question (needs a person's decision, affects the API shape):**
  should the cap be a single fixed constant inside `Slugify`, a
  package-level default that's still a plain `Slugify(s string) string`
  call, or a parameter on every call (e.g. an additional `n int`, the
  same shape as `Truncate(s string, n int) string` from issue #7)? A
  parameter changes `Slugify`'s existing signature and every call site;
  a constant does not but can't vary by consumer. This is ultimately an
  `/architect-design` decision, but it depends on the product question
  of whether different consumers need different limits — no evidence of
  a second consumer with different needs exists yet, so the recommended
  default is one global limit (no signature change) until one does.
- **Assumption (unvalidated):** the cap should be measured on the final
  `Slugify` output, i.e. applied after any separator collapsing/trimming
  (issue #3) and after any transliteration (issue #5's P0), not on an
  intermediate form — so the cap always reflects what the caller actually
  receives. Flagged here so `/architect-design` doesn't have to
  rediscover it, not decided as a product trade-off.

## Hand-off note

Under normal operation this PRD would now hand off to delivery: one
GitHub issue per P0 story, labeled `product`, with the acceptance
criteria above in the body, so `/pm-sprint-delivery` can pull it into an
umbrella — but only once the open questions above are answered, since the
story's acceptance criteria can't be made concrete (no cap value, no
settled API shape) until then. No such issue was filed because this run's
platform rules route output into this document instead of creating
issues/PRs/comments.

> PRD written to `docs/product/slugify-length-cap.md`: 1 P0 story (cap
> Slugify's output length, cut at a separator boundary, trim any
> resulting trailing dash), 2 deferred (script-aware limits, per-consumer
> configurable cap).
> Success metric: % of capped slugs ending in a dangling `-`, target 0%;
> real slug-length distribution to be gathered before the cap value is
> fixed.
> Next: a person answers the two open questions above (the cap value and
> the API shape), then `/architect-design slugify-length-cap` for the
> implementation, then `/pm-sprint-delivery` to schedule the P0 story.
