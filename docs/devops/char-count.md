# DevOps plan: Character count helper for the `text` package

**Date:** 2026-10-15
**Owner:** unassigned
**Stage:** devops (flow `d358b357`, stage 8 of 8)
**Skill:** `devops-deploy`, adapted for this stage to a documents-only
output — nothing in this pass was deployed, no container was created or
torn down, no issue/PR/comment was opened, and no code, test, or secret was
touched. Everything the skill would normally do or post is recorded in this
plan instead.

## Pre-flight: what exactly would be deploying

Per `docs/architecture/char-count.md` and `docs/sprint/char-count.md`, the
change this flow produced is:

- `text/char_count.go` — one new exported function, `CharCount(s string)
  int`, in the existing `text` package
- `text/char_count_test.go` — its table-driven tests
- `README.md` — one line added to the file/purpose table

There is no new service, no container, no network boundary, no database,
and no compose file anywhere in this diff. The architecture doc states this
explicitly in its "Deviations from the default stack" table: nothing here
is deployable, it is library code consumed at compile time by whatever
already imports `text`.

This repository is also, per its own `README.md`, a fixture used to
exercise the issue-to-PR flow — "Nothing here is a product" — and has no
standing `demo` or `prod` Containarium backend, no `compose` file, and no
route exposed anywhere in the tree. There is no backend to run
`list_backends` / `get_backend` against, and no container to run
`list_containers` / `create_container` against.

**Conclusion:** this stage has nothing to sync, run, or expose. The
`devops-deploy` steps that assume a running backend (2–5, 7, 8) do not
apply, for the same reason `docs/qa/char-count.md` found the batched
verification pass (step 7 of `qa-sprint-verify`) inapplicable: the sprint's
own **Verification queue** is empty, because every acceptance criterion is
provable by `go vet ./...` and `go test ./...` — CI, not a deployed
environment.

## Deployment plan (what ships this, and how)

Since the unit of deployment for this change is a git commit merged to
`main`, not a container sync, the plan is:

1. **Gate on CI, not on a container.** The pull request carrying
   `text/char_count.go`, `text/char_count_test.go`, and the `README.md`
   line must show `go vet ./...` and `go test ./...` green in
   `.github/workflows/ci.yml` before merge. A commit that has not passed
   those checks does not merge, by the same rule the skill applies to a
   container deploy ("a commit that never passed tests does not get
   deployed").
2. **Pre-merge check.** A human reviewer confirms the diff touches only the
   three files named above and does not change `text/slug.go` or
   `text/slug_test.go` — `docs/qa/char-count.md` already confirmed this by
   `git diff` on the flow branch; the reviewer re-confirms it on the actual
   PR diff before merging, since that is the version that ships.
3. **Merge to `main`.** This is the entire "deploy" — once merged, every
   future importer of the `text` package picks up `CharCount` the next time
   it builds. No sync step, no `compose_enable`, no `expose_port`: there is
   no remote process to start and no route to verify, because this package
   is a library, not a running service.
4. **Post-merge check.** Confirm `main`'s CI is still green immediately
   after the merge commit lands (the same `go vet` / `go test` gate, now
   run against `main` instead of the PR branch). A red `main` after merge
   is the equivalent of a failed health check in a container deploy and is
   treated the same way: fix forward or revert, not left red.
5. **No secrets, no environment, no route.** Nothing in this change reads a
   config value, a credential, or an environment variable, so the
   platform's secret store is not involved at any step.

## Rollback

If `CharCount` turns out to be wrong after merge (for example, a case the
table tests missed surfaces a bug), rollback is `git revert` of the merge
commit on `main`, followed by the same CI gate on the revert commit before
it lands. There is no container to redeploy to a previous version and no
data migration to undo, since this change touches no state outside the
source tree.

## Who must act

- **A human reviewer** merges the pull request once CI is green and the
  diff scope is confirmed (step 2 above) — this plan does not merge
  anything itself.
- **No DevOps/Containarium action is required**: there is no backend,
  container, or route for this repository to create, sync, or tear down,
  and no sprint-umbrella verification window to open, since
  `docs/sprint/char-count.md`'s Verification queue is empty and names CI as
  sufficient proof.

## Guardrails observed in this stage

- No credential, token, key, password, or other secret value appears
  anywhere in this document.
- No command in this document is meant to be run against a production
  system; every step above is a description of what a reviewer or CI
  checks, not an executable instruction against a live backend.
- No container was created, synced, or deleted; no umbrella comment was
  posted; no code or test file was changed. This document is the only
  artifact this stage produces.
