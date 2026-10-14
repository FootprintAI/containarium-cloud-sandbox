# Design: ReadingTime helper for the text package

**Date:** 2026-10-14
**Status:** proposed
**Stack:** N/A (default stack) — pure Go standard-library addition to the
existing `text` package; no service, network contract, container, or
frontend boundary is introduced. See Deviations below.

## Problem

`docs/product/reading-time.md` (the product stage of this flow) scopes a
single function: `text.ReadingTime(s string) time.Duration`, estimating
reading time at a fixed 200 words/minute, ceiling-rounded to a whole
minute, with the one-minute floor and zero-for-empty/whitespace-only
behavior the PRD specifies. This doc designs that one function and its
test strategy — nothing else is in scope.

## Design

One component, already built once (`text/slug.go`, `Slugify`) and
designed similarly before (issue #7's `Truncate`): a pure function added
to the existing `text` package, with no new package, service, or
dependency.

```
text/
  reading_time.go       // new: ReadingTime
  reading_time_test.go  // new: table-driven tests
  slug.go                // existing, unchanged
  slug_test.go           // existing, unchanged
```

No diagram: a single function, no data flow across a boundary, no
process to deploy.

**Function shape** (implementation detail, included here so the
implementing task doesn't have to redesign it):

```go
package text

import (
	"strings"
	"time"
)

const wordsPerMinute = 200

// ReadingTime estimates how long s takes to read at 200 words per
// minute, rounded up to a whole minute. Empty or whitespace-only text
// returns 0; any other non-empty text returns at least one minute.
func ReadingTime(s string) time.Duration {
	words := len(strings.Fields(s))
	if words == 0 {
		return 0
	}
	minutes := (words + wordsPerMinute - 1) / wordsPerMinute
	return time.Duration(minutes) * time.Minute
}
```

`strings.Fields` already treats any run of whitespace as a single
separator and ignores leading/trailing whitespace, which is exactly the
PRD's "whitespace-separated tokens" word-counting rule — no extra
trimming or regex needed. `(words + wordsPerMinute - 1) / wordsPerMinute`
is integer ceiling division: for any `words >= 1` it is already `>= 1`,
so the PRD's "one-minute floor" falls out of the arithmetic rather than
needing a separate clamp — worth calling out explicitly so the
implementing task doesn't add a redundant `if minutes < 1` branch.

## Language choices

| Component | Language | Why this one | Type gate in CI |
|-----------|----------|---------------|-----------------|
| `text.ReadingTime` | Go | Matches the language of the existing `text` package (`Slugify`, and `Truncate` from issue #7) — this is one more function in that package, not a new component choosing its own language | `go vet` / `go build` (existing `ci.yml` already runs this on the package) |

No second or third language is introduced; the "deviation needs a
reason" rule doesn't apply here because nothing deviates.

## Contracts

None. `ReadingTime` is an in-process pure function call within a single
Go binary's dependency graph — there is no network boundary, no
serialized wire format, and therefore no `.proto`, REST mapping, or
generated client to define. (Contrast with this skill's default
stack, which assumes a service boundary; see Deviations.)

## Test strategy

Table-driven tests in `text/reading_time_test.go`, next to the function,
matching the existing convention in `slug_test.go`. Every row traces to
an acceptance criterion in `docs/product/reading-time.md`:

| Case | Input shape | Expected | Acceptance criterion covered |
|------|-------------|----------|-------------------------------|
| Empty string | `""` | `0` | "`ReadingTime("")` returns `0`" |
| Whitespace-only | `"   \t\n  "` | `0` | "whitespace-only string ... returns `0`" |
| Single short word | `"Hello"` | `1 * time.Minute` | one-minute floor |
| Few words, still under 200 | `"Hello there friend"` | `1 * time.Minute` | one-minute floor |
| Exactly 200 words | `strings.Repeat("w ", 200)` (200 whitespace-separated tokens) | `1 * time.Minute` | ceiling of an exact multiple stays whole, not bumped up |
| Exactly 400 words | `strings.Repeat("w ", 400)` | `2 * time.Minute` | same, at a second multiple (the PRD's own worked example) |
| 450 words (a remainder) | `strings.Repeat("w ", 450)` | `3 * time.Minute` | remainder rounds up (the PRD's own worked example: 2.25 → 3) |

Using `strings.Repeat("w ", n)` to build the longer fixtures keeps the
test deterministic and readable (it's immediately obvious the fixture
has exactly `n` words) without pulling in real prose or a corpus file.

No contract tests (no boundary) and no integration/e2e slice (no system
to assemble) — this is the whole test surface. No mocks: `ReadingTime`
has no dependency to fake (no clock, no I/O), so every test runs
real-value-in, real-value-out, including `time.Duration` equality
checks.

**Not covered by this test strategy, by design:** the PRD leaves CJK/
non-space-delimited word counting and configurable reading speed as open
product questions, not yet decided. No test asserts behavior for
no-whitespace scripts or a non-200 rate, because asserting one now would
silently pick an answer to a question the PRD explicitly left open.

## Deviations from the default stack

- **No protobuf/gRPC contract, no Docker packaging, no Next.js/
  TypeScript layer.** The default stack in this skill assumes a
  deployable service or a user-facing frontend. This change is a single
  pure function inside an existing Go library package with no network
  boundary and nothing to run or ship as a container — there is no
  service for a contract or a Dockerfile to describe. Containment
  boundary: `text/reading_time.go` only; nothing else in the repo's
  deployment shape changes.
- Everything else (Go as the language, table-driven tests, `go vet` as
  the CI type gate) already matches the default stack; no further
  deviation.

## Rejected alternatives

- **Regex-based word splitting** (e.g. `regexp.MustCompile(\`\s+\`)`
  plus manual trim) instead of `strings.Fields`. Rejected: `strings.Fields`
  is the standard-library function built exactly for this — split on
  runs of whitespace, ignore leading/trailing whitespace — with no
  compiled pattern, no extra import beyond `strings` (already a
  dependency of this package via `slug.go`), and no behavior to get
  subtly wrong relative to a hand-written regex.
- **Floating-point division plus `math.Ceil`** (`int(math.Ceil(float64(words)/200))`)
  instead of integer ceiling arithmetic. Rejected: integer ceiling
  division (`(words + wordsPerMinute - 1) / wordsPerMinute`) gives the
  exact same result for these inputs without a `math` import or any
  floating-point rounding edge case to reason about, and keeps the
  function consistent with this package's existing "standard library,
  minimal footprint" style (`slug.go` imports only `strings`).
