# Design: WordCount helper for the text package

**Date:** 2026-10-15
**Status:** proposed
**Stack:** N/A (default stack) — pure Go standard-library addition to
the existing `text` package; no service, network contract, container,
or frontend boundary is introduced. See Deviations below.

## Problem

`docs/product/word-count.md` (the product stage of this flow) scopes a
single function: `text.WordCount(s string) int`, returning the number
of maximal runs of non-space characters in `s`. This doc designs that
one function and its test strategy — nothing else is in scope.

## Design

One component: a pure function added to the existing `text` package,
alongside its existing `Slugify` function, with no new package, service,
or dependency.

```
text/
  word_count.go       // new: WordCount
  word_count_test.go  // new: table-driven tests
  slug.go              // existing, unchanged
  slug_test.go         // existing, unchanged
```

No diagram: a single function, no data flow across a boundary, no
process to deploy.

**Function shape** (implementation detail, included here so the
implementing task doesn't have to redesign it):

```go
package text

import "strings"

// WordCount returns the number of maximal runs of non-space characters
// in s. An empty or whitespace-only s returns 0.
func WordCount(s string) int {
	return len(strings.Fields(s))
}
```

`strings.Fields` already splits `s` on runs of whitespace and discards
any empty results from leading/trailing whitespace — its return length
*is* the count the PRD asks for, exactly. There is no additional logic
to write beyond calling it.

## Language choices

| Component | Language | Why this one | Type gate in CI |
|-----------|----------|---------------|-----------------|
| `text.WordCount` | Go | Matches the language of the existing `text` package (`Slugify`) — one more function in that package, not a new component choosing its own language | `go vet` / `go build` (existing `ci.yml` already runs this on the package) |

No second or third language is introduced; the "deviation needs a
reason" rule doesn't apply because nothing deviates.

## Contracts

None. `WordCount` is an in-process pure function call within a single Go
binary's dependency graph — there is no network boundary, no serialized
wire format, and therefore no `.proto`, REST mapping, or generated
client to define.

## Test strategy

Table-driven tests in `text/word_count_test.go`, next to the function,
matching the existing convention in `slug_test.go`. Every row traces to
an acceptance criterion in `docs/product/word-count.md`:

| Case | Input | Expected | Acceptance criterion covered |
|------|-------|----------|-------------------------------|
| Empty string | `""` | `0` | "`WordCount("")` returns `0`" |
| Whitespace-only | `"   \t\n  "` | `0` | "whitespace-only string ... returns `0`" |
| Single word | `"Hello"` | `1` | "a single-word string returns `1`" |
| Multiple words, single spaces | `"Hello there friend"` | `3` | "a string of N words ... returns `N`" |
| Multiple words, mixed whitespace | `"Hello\tthere\n\nfriend"` (tab and blank line as separators) | `3` | same criterion, covering tabs/newlines/multiple separators |
| Punctuation attached to a word | `"word, word."` | `2` | "a run of non-space characters that includes punctuation ... counts as one word" |

No contract tests (no boundary) and no integration/e2e slice (no system
to assemble) — this is the whole test surface. No mocks: `WordCount` has
no dependency to fake (no clock, no I/O), so every test runs
real-value-in, real-value-out.

## Deviations from the default stack

- **No protobuf/gRPC contract, no Docker packaging, no Next.js/
  TypeScript layer.** The default stack in this skill assumes a
  deployable service or a user-facing frontend. This change is a single
  pure function inside an existing Go library package with no network
  boundary and nothing to run or ship as a container — there is no
  service for a contract or a Dockerfile to describe. Containment
  boundary: `text/word_count.go` only; nothing else in the repo's
  deployment shape changes.
- Everything else (Go as the language, table-driven tests, `go vet` as
  the CI type gate) already matches the default stack; no further
  deviation.

## Rejected alternatives

- **A hand-written byte-scanning loop** that tracks transitions between
  "in a word" and "in whitespace" instead of calling `strings.Fields`.
  Rejected: it would reimplement exactly what `strings.Fields` already
  does correctly — split on any run of Unicode whitespace, discard empty
  results from leading/trailing whitespace — with more code, no
  behavioral difference, and more surface area to get subtly wrong.
- **A shared internal word-splitting helper** reused by `WordCount` and
  any future helper that needs a word count (the PRD's own "Later
  phases" note raises this possibility). Rejected for this design: no
  second consumer exists yet in this flow, so introducing shared
  internal plumbing now would be structure built ahead of a second use
  that may never arrive — exactly what the PRD defers until it does.
