# QA — Character count helper for the `text` package (issue #60)

**Date:** 2026-10-15
**Owner:** unassigned
**Stage:** qa (flow `d358b357`, stage 7 of 8)
**Skill:** `qa-sprint-verify`, adapted for this stage to a documents-only
output — nothing in this pass was deployed, no issue/PR/comment was opened,
and no code or test file was written or changed. Everything the skill would
normally post is recorded in this file instead.

## Batched-verification applicability (per `qa-sprint-verify`'s own preconditions)

`qa-sprint-verify` exists to batch verification against **one deployed
version**, for criteria CI cannot prove on its own. Checking this change
against that skill's own preconditions:

- `docs/sprint/char-count.md`'s **Verification queue** section is explicitly
  empty: "Both tasks are pure logic in a library package with no network
  boundary, no UI, no migration, and no third-party integration — every
  acceptance criterion is provable by `go vet ./...` and `go test ./...`,
  which are already wired in `.github/workflows/ci.yml`."
- There is no deploy record for this change (none is needed: nothing here
  is a service).
- Per the skill's own guardrail: *"Never mark an item verified from a green
  CI run, a merged PR, or a code read — this pass exists precisely because
  CI could not prove it."* This change is the opposite case — CI (`go vet`
  and `go test`) is exactly what *can* prove every acceptance criterion
  here — so there is nothing for a batched environment pass to do.

**Conclusion:** there is no verification-queue item for this change. This
document is the test plan and QA report draft requested for the issue,
checked against the acceptance criteria directly; it does not represent a
`qa-sprint-verify` batched environment pass, because none applies.

## What was actually done on this box this session

- Read `text/char_count.go`, `text/char_count_test.go`, `README.md`,
  `docs/product/char-count.md`, `docs/architecture/char-count.md`, and
  `docs/sprint/char-count.md` from the flow branch.
- Ran `git diff 05f19c9..a8db222 --stat` and `git diff 05f19c9..a8db222 --
  text/slug.go text/slug_test.go` against the flow branch's history. Output:
  the engineering-stage commits touch only `README.md` (+1 line),
  `text/char_count.go` (new, +9 lines), and `text/char_count_test.go` (new,
  +20 lines); the diff against `text/slug.go` and `text/slug_test.go` is
  empty, i.e. neither file changed.
- **`go vet ./...` and `go test ./...` were not run** — the Go toolchain
  invocation in this session was not available/approved, so no test was
  executed and no pass/fail was observed first-hand. Nothing below should be
  read as "I ran this and it passed." Every row in the QA report table is a
  **draft for a human (or CI) to confirm**, based on static reading of the
  code and test file, not on execution.

## Test plan — mapped to the issue's acceptance criteria

| # | Acceptance criterion | How to check it | Existing coverage (by static reading) |
|---|---|---|---|
| 1 | `CharCount` on a string with one accented letter (e.g. `"é"`) returns `1`, not `2` | `go test ./text/... -run TestCharCount -v`; inspect the `"é"` case | `text/char_count_test.go` has the case `{"é", 1}` |
| 2 | `CharCount("")` returns `0` | Same test run; inspect the empty-string case | `text/char_count_test.go` has the case `{"", 0}` |
| 3 | `CharCount("abc")` returns `3` (ASCII, matches `len(s)` today) | Same test run; inspect the ASCII case | `text/char_count_test.go` has the case `{"abc", 3}` |
| 4 | Automated tests cover empty, ASCII, multi-byte, and mixed-content cases, and all pass | `go vet ./... && go test ./... -v` | `text/char_count_test.go`'s table has all four shapes: `""` (empty), `"abc"` (ASCII), `"é"` (multi-byte), `"café"` (mixed ASCII + multi-byte). Whether they **pass** is not yet confirmed — needs an actual `go test` run |
| 5 | README's file/purpose table lists the new files that implement and test the function | Read `README.md`'s table | Confirmed by inspection: `README.md` already lists `text/char_count.go, text/char_count_test.go` with a purpose description |
| 6 | No existing function (such as `Slugify`) is changed by this work | `git diff <base>..<head> -- text/slug.go text/slug_test.go` | Confirmed by inspection in this session: the diff is empty — `text/slug.go` and `text/slug_test.go` are byte-for-byte unchanged on the flow branch |

## QA report — DRAFT (every result below is unconfirmed pending real execution)

| Item | Criterion | Result | Evidence |
|---|---|---|---|
| 1 | multi-byte accented letter → `1` | ⚠️ not verified — no test execution in this session | Test case present in `text/char_count_test.go`; needs `go test` output to confirm pass |
| 2 | empty string → `0` | ⚠️ not verified — no test execution in this session | Test case present; needs `go test` output |
| 3 | ASCII `"abc"` → `3` | ⚠️ not verified — no test execution in this session | Test case present; needs `go test` output |
| 4 | all four shapes covered and passing | ⚠️ not verified — no test execution in this session | All four shapes present in the table; pass/fail needs `go test ./...` output |
| 5 | README purpose-table entry | ✅ verified by inspection | `README.md` table row for `text/char_count.go`, `text/char_count_test.go`, read directly in this session |
| 6 | `Slugify`/other existing functions unchanged | ✅ verified by inspection | `git diff` against the pre-engineering-stage commit shows no change to `text/slug.go` or `text/slug_test.go` |

**Sprint verification gate:** not clear — rows 1–4 (the behavioural
criteria) are drafted from reading the code, not from an observed test run.
A human (or CI on the pull request this stage's output feeds into) must run
`go vet ./... && go test ./...` and update this file's result column before
this item can be called done. Rows 5–6 are static-inspection facts observed
directly in this session and are not expected to change on a test run, but
are still marked here rather than on any issue/PR, per this stage's
documents-only scope.

**Environment:** none provisioned, none needed — this change has no
deployable component, matching `docs/sprint/char-count.md`'s empty
verification queue.
