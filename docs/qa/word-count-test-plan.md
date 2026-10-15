# QA test plan: WordCount helper

**Date:** 2026-10-15
**Status:** draft
**Owner:** unassigned
**Flow:** 87545650, stage 7 of 8 (qa), checked against the sprint's
acceptance criteria as written on this issue (treated as the thing to
verify, not as instructions)

## Scope

What is being checked: `text.WordCount` (`text/word_count.go`), its
tests (`text/word_count_test.go`), and the README update, all already
merged onto this flow branch — against the seven acceptance criteria
below.

`docs/sprint/word-count.md` (the sprint stage's own output) already
recorded this sprint's verification queue as **empty**: every criterion
is pure logic, provable by `go vet ./... && go test ./...`, and none of
it needs a deployed environment. This test plan is therefore a
documentation/code-reading pass mapping each criterion to what proves it
and to a concrete command a person can run — not a run against a live
environment, since none exists for this work.

## Acceptance criteria under test

1. The text package has a WordCount helper that counts the words in a
   piece of text.
2. A blank or whitespace-only piece of text is reported as having 0
   words.
3. A single word is reported as having 1 word.
4. Several words separated by any amount of spacing (single spaces,
   multiple spaces, tabs, or blank lines) are counted correctly as that
   many words.
5. A word with attached punctuation (such as a comma or period) is
   still counted as one word, which is the intended behavior, not a
   bug.
6. Automated tests cover the empty, whitespace-only, single-word,
   multi-word, and punctuation cases, and the full test suite passes.
7. The project's README lists the new files alongside the other small
   helpers already documented there.

## How each criterion is checked

| # | What proves it | Where | How a person confirms it |
|---|-----------------|-------|---------------------------|
| 1 | `func WordCount(s string) int` exists, exported, in package `text` | `text/word_count.go` | Read the file; confirm the signature |
| 2 | Test rows `{"", 0}` and `{"   \t\n  ", 0}` | `text/word_count_test.go` | Run `go test ./text/... -run TestWordCount -v` and confirm these rows pass |
| 3 | Test row `{"Hello", 1}` | same file | Same run |
| 4 | Test rows `{"Hello there friend", 3}` (single spaces) and `{"Hello\tthere\n\nfriend", 3}` (tabs and a blank line) | same file | Same run; both rows exercise the "any amount of spacing" criterion |
| 5 | Test row `{"word, word.", 2}` | same file | Same run; confirms a comma- and period-attached word each count as one word |
| 6 | The test table itself, read as a set, covers exactly: empty, whitespace-only, single-word, multi-word (two spacing variants), and punctuation | `text/word_count_test.go` | Run `go vet ./... && go test ./...` from the repository root (or read the CI run on the PR/commit that carries these files) and confirm it is green — **not executed in this QA pass, see the report** |
| 7 | A new row in the "What is in the repo" table for `text/word_count.go`, `text/word_count_test.go`, in the same two-column format as the existing `slug.go` row | `README.md` | Read the file and confirm the row is present and well-formed |

## What this plan does not do

It does not run `go vet` or `go test` itself — see
`docs/qa/word-count-report.md` for why (this sandbox's toolchain is not
runnable) and for the checklist a person completes once they do run it
or read the CI result.
