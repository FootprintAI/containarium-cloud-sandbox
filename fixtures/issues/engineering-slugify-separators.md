---
title: Slugify should collapse repeated separators and trim the ends
labels: scope:engineering
---

`text.Slugify` turns every non-alphanumeric character into its own `-`, so
punctuation and spaces produce noisy slugs:

| Input | Today | Expected |
| --- | --- | --- |
| `Hello,  World!!` | `hello---world--` | `hello-world` |
| `  Padded  ` | `--padded--` | `padded` |
| `--x--` | `--x--` | `x` |

Make `Slugify` collapse each run of separators into a single `-` and trim a
leading or trailing `-`. Keep the existing cases in `text/slug_test.go`
passing, and add test cases for the three rows above.

Acceptance:

- [ ] the three rows above produce the expected slug
- [ ] existing cases still pass (`go test ./...`)
- [ ] only `text/slug.go` and `text/slug_test.go` change
