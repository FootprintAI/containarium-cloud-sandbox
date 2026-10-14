# Deployment plan (DRAFT): ReadingTime helper

**Date:** 2026-10-15
**Status:** draft — not executed
**Owner:** unassigned
**Flow:** f16c5d49, stage 6 of 6 (devops)

## What this plan is

A plan only. Per this stage's rules, nothing in it is deployed, merged,
or run by this pass — every action below names a human operator who
carries it out. The platform does not deploy this, and neither does this
pass.

## What is actually being shipped

`text.ReadingTime` (`text/reading_time.go`), its tests
(`text/reading_time_test.go`), and one new README row — a pure Go
library addition with no running service, no container, no network
route, and no configuration or secret of any kind. Confirmed by reading
this repository's full file tree on this flow branch: there is no
Dockerfile, no compose file, and no server entry point anywhere in it;
`go.mod` declares a library module, not a deployable.

## Why this skill's usual deploy steps don't apply as written

`/devops-deploy`'s default model targets a running service on
Containarium: sync committed state to a container, run it, verify a
health route, then open a verification window for the deployed commit.
There is no container to sync to and no route to verify here. This is
also consistent with the earlier stages of this same flow:
`docs/sprint/reading-time.md` already recorded this sprint's
verification queue as empty, because every acceptance criterion is
provable by `go vet`/`go test` with no deployed environment needed. The
plan below keeps the same principles this skill cares about —
committed state only, a human confirms before anything ships, rollback
stays boring and git-native — applied to what "shipping" actually means
for a library change: merging reviewed, green commits into `main`.

## Deployment plan

### Pre-flight (before anything merges)

1. Confirm the branch(es) carrying this flow's engineering-stage changes
   are based on the current `main`, and that `main`'s own CI
   (`go vet ./... && go test ./...`, per
   `.github/workflows/ci.yml`) was green at the point they branched from.
2. Confirm CI is green on each of those branches' own pull request(s) —
   a commit that hasn't passed `go vet`/`go test` does not merge, per
   this repository's own rule that `main` must stay green.
3. Confirm `docs/qa/reading-time-report.md`'s checklist has actually
   been completed by a person against a real `go vet`/`go test` run or
   the CI result. It was left as an explicit, unconfirmed draft by the
   QA stage — merging ahead of that would ship work nobody has actually
   checked.

### Steps, in order (who must act)

1. **A repository maintainer** reviews the pull request(s) carrying the
   engineering-stage changes, confirms CI is green on them, and
   approves.
2. **The same maintainer** completes (or confirms someone already
   completed) the QA checklist in `docs/qa/reading-time-report.md`
   before merging — not this pass, and not automatically.
3. **The maintainer** merges the approved pull request(s) into `main`
   through the normal GitHub merge flow.
4. **The maintainer** confirms `main`'s own CI run, triggered by the
   merge (`.github/workflows/ci.yml`'s `push: branches: [main]`
   trigger), comes back green immediately after. A red run on `main`
   means the merge broke the baseline the README promises stays green,
   and is treated as a failed deploy.
5. No service restart, container sync, route check, or health check is
   needed or possible here — there is nothing running to restart or
   check.
6. No secret is read, written, or required anywhere in this plan; it
   does not touch the platform's secret store, and no operator needs to
   supply one.

### Verification window

Not opened. The sprint's verification queue
(`docs/sprint/reading-time.md`) is already recorded empty — every
acceptance criterion is CI-provable, so there is nothing to hand off to
a batched verification pass against a deployed environment, and no
container is created or torn down for this flow.

### Rollback

If `main`'s CI goes red after the merge, or a defect surfaces later: **a
maintainer** reverts the merge commit on `main` through the normal
git/GitHub revert flow and confirms CI is green again on the reverted
state. No data changed shape and nothing is running, so no
backup/restore step applies — this is the git-native rollback this
skill favors, at its smallest possible scope.

## Who must act (summary)

- **A human repository maintainer:** reviews the pull request(s),
  confirms CI, completes the QA checklist, performs the merge, and would
  perform any revert.
- **This pass / the platform:** writes this plan only. No merge, no
  revert, and no container action is taken here.
