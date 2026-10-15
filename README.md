# containarium-cloud-sandbox

A small, public fixture repository for exercising Containarium Cloud's
issue-to-pull-request flow end to end: a user connects a GitHub repo, files an
issue, and a hosted agent opens a pull request. Nothing here is a product; it
exists so the flow can be run against something real and then cleaned up.

## What is in the repo

| Path | Purpose |
| --- | --- |
| `text/slug.go`, `text/slug_test.go` | One tiny function with a green test suite. It has one known, deliberate gap (see the issue fixtures), small enough for a single focused change. |
| `text/char_count.go`, `text/char_count_test.go` | `CharCount` returns the number of Unicode code points (runes) in a string, not its byte length, with a table-driven test suite covering empty, ASCII, and multi-byte input. |
| `fixtures/issues/engineering-slugify-separators.md` | A solvable issue: a clear gap, a table of expected behaviour, acceptance criteria. Expect a small PR touching `text/slug.go` and `text/slug_test.go`. |
| `fixtures/issues/unresolvable-question.md` | A product question, not a code change. Expect **no edits** and a stated reason. |
| `.github/labels.json` | The label set the tracker flow uses (`scope:*` routes an issue to a skill, `agent:*` is run state). |
| `.github/workflows/ci.yml` | `go vet` and `go test` on every PR, so an agent's PR gets a real check. |

The baseline must stay green: `go vet ./... && go test ./...` passes on `main`.
The slug gap is intentionally **not** fixed on `main`; the agent is meant to fix
it.

## Test plan (the agent-on-an-issue steps)

1. **Labels.** Create the labels once:
   `jq -c '.[]' .github/labels.json | while read -r l; do gh label create "$(jq -r .name <<<"$l")" --color "$(jq -r .color <<<"$l")" --description "$(jq -r .description <<<"$l")" --force; done`
2. **Connect** the repo to the Containarium Cloud organisation under test (GitHub App
   installation) and enable the issue tracker for it. Route
   `scope:engineering` to the implementing skill.
3. **File the solvable issue** from `fixtures/issues/engineering-slugify-separators.md`
   (title and body as written, label `scope:engineering`). Creating the issue alone
   starts nothing; the label is the trigger.
4. **Expect, within about two dispatch intervals:** the issue gains `agent:queued`,
   then `agent:running`; the agent comments with its run identity; a pull request
   appears whose diff touches only `text/slug.go` and `text/slug_test.go`, with the
   three table rows added as test cases; CI on that PR is green; the issue ends
   `agent:done`.
5. **File the unresolvable issue** from `fixtures/issues/unresolvable-question.md`
   with the same label. Expect a visible terminal state with **no pull request**
   and a comment explaining why nothing was changed.
6. **Credential hygiene.** The agent's box must hold no GitHub credential; the
   pull request is opened by the platform from the agent's output.
7. **Clean up.** Close the issues and the pull request, delete the agent's branch,
   remove the label state (`agent:*`) if left behind, and disconnect the repo.
   Do not merge agent pull requests into `main`: the gap has to stay open for the
   next run.

## Rules for this repo

- Public. Never put hostnames, account names, tokens or organisation names of any
  deployment in an issue, a comment or a file.
- The `main` branch is a fixture, not a product: keep it small and green.
