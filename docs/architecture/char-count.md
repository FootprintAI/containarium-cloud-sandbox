# Design: Character count helper for the `text` package

**Date:** 2026-10-15
**Status:** proposed
**Owner:** unassigned
**Stack:** Go (standard library only) — no protobuf/gRPC, no Docker, no
Next.js/TypeScript, no Python; see Deviations below for why none of those
apply here.

## Problem

Per `docs/product/char-count.md` (issue #50), Go's `len(s)` counts bytes, so
callers measuring a string's visible length get the wrong answer for any
multi-byte rune (e.g. an accented letter). The `text` package needs a
`CharCount(s string) int` function that returns the rune count instead, so
that "characters" means Unicode code points consistently with how
`text.Slugify` already walks the string rune by rune.

## Design

This is a single pure function added to the existing `text` package — no new
service, no deployable, no network boundary. It sits next to `Slugify` in
the same package and has the same shape: in-process, stateless, no side
effects, no I/O.

- **Component:** `text.CharCount` — a new exported function in a new file,
  `text/char_count.go`, in the existing `text` package.
- **Behaviour:** `CharCount(s string) int` returns the number of Unicode
  code points in `s`. Implement it with the standard library's
  `unicode/utf8.RuneCountInString`, which does exactly this in one pass with
  no allocation — there is no reason to hand-roll a `range`-based counting
  loop (that duplicates what the stdlib already does correctly and is
  exactly the kind of logic most likely to pick up an off-by-one bug).
- **Data flow:** caller → `CharCount(s)` → `utf8.RuneCountInString(s)` →
  `int`. No failure path exists: every Go string is valid input (invalid
  UTF-8 bytes are counted as one `utf8.RuneError` rune each, matching
  `range`'s own behaviour, so this stays consistent with `Slugify`).
- No diagram: one function, one call, no topology to show.

## Language choices

| Component | Language | Why this one | Type gate in CI |
|-----------|----------|---------------|-----------------|
| `text.CharCount` | Go | Extends the existing `text` package, which is already Go; a one-line stdlib call needs no other runtime | `go vet ./...`, `go build ./...` |

## Contracts

None. This is an in-process function call within a single Go package, not a
service boundary — there is no wire format, no client/server pair, and
nothing for a second language to generate a client against. The Go function
signature `CharCount(s string) int` is itself the only contract, enforced by
the Go compiler.

## Test strategy

- **Unit tests (table-driven, `text/char_count_test.go`):**
  - empty string → `0`
  - ASCII-only string (e.g. `"abc"`) → same value as `len(s)` (`3`)
  - a string containing one multi-byte rune (e.g. `"é"`) → `1`, not `2`
  - a string mixing ASCII and multi-byte runes, to prove the count is per
    rune across the whole string, not just at a boundary
- **Contract tests:** none — no cross-component or cross-language boundary
  exists for this change.
- **Integration/e2e:** none needed beyond the unit tests — `CharCount` has
  no I/O or external dependency to exercise at a larger scope.
- **Mocking:** none. Pure function, no collaborators to fake or stub.
- **CI gate:** `go vet ./...` and `go test ./...` (already wired in
  `.github/workflows/ci.yml`) cover both the type check and the test suite
  for this addition; no new CI lane is needed.

## Deviations from the default stack

| Default | Why it doesn't apply here |
|---------|---------------------------|
| Protobuf/gRPC + grpc-gateway contracts | No service and no network boundary exist in this change — it's a single in-process function in an existing library package. Introducing a contract layer for a function call would add a build/generation step with nothing on the other side of it to type. |
| Docker | Nothing here is deployable; it's library code consumed at compile time by whatever already imports `text`. |
| Next.js/React/TypeScript, Python | No frontend and no ML/data workload is involved. The only sanctioned language already in play, Go, is also the correct one for this change. |

All deviations are contained entirely within the scope of this one change
(`text/char_count.go` and its test file) and do not set a precedent for the
rest of the repo, which keeps the Docker/protobuf/gRPC stack for anything
that is an actual service.

## Rejected alternatives

- **Hand-rolled rune-counting loop** (`for range s { n++ }`, mirroring
  `Slugify`'s own loop style) — rejected in favor of
  `utf8.RuneCountInString`, which is the standard library doing the same
  thing with no risk of an off-by-one or allocation bug in hand-written
  code that would need its own test just to prove it matches the stdlib.
- **Grapheme-cluster-aware counting** (treating combining-mark sequences or
  modifier emoji as one visual character via a library such as
  `rivo/uniseg`) — rejected because the PRD explicitly scopes this to a
  rune count and puts grapheme segmentation out of scope; pulling in a
  dependency for behaviour nobody asked for would be scope creep.
