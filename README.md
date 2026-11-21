# go-verdict

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

The fleet's typed epistemic status: five states, one worst-of roll-up, one render grammar — extracted once from the ~14 dashboards and gate batteries that hand-rolled it (2026-11-21 domain-language census).

```go
import "github.com/larsartmann/go-verdict"

v := verdict.WorstOf(gateFreshness, gateStub, gateCitations)
fmt.Printf("%s battery: %s\n", v.Icon(), v) // "✓ battery: PASS" / "✗ battery: FAIL" / ...
os.Exit(v.ExitCode())                        // only FAIL exits 1
```

## The states, and what each promises

| State     | Promise                                              | Fails a roll-up? |
| --------- | ---------------------------------------------------- | ---------------- |
| `PASS`    | ran, and the claim it tests holds                    | no               |
| `FAIL`    | ran, and the claim is false — the only failing state | **yes**          |
| `WARN`    | ran, holds, but carries debt the reader must see     | no               |
| `SKIP`    | deliberately not evaluated here (exempt/filtered)    | no (never votes) |
| `UNKNOWN` | cannot say — never ran, timed out, inconclusive      | no               |

## Roll-up (`WorstOf`)

`Fail > Warn > Unknown > Pass > Skip`. Skips never vote; empty or all-Skip input rolls up to `UNKNOWN` (nothing was said — an honest roll-up does not invent a Pass). An `UNKNOWN` poisons an otherwise all-Pass roll-up: unverified is not verified.

## Parse

`Parse(" fail ")` → `FAIL`; any unrecognized word (including `""` and near-misses like `"passed"`) → `UNKNOWN`. For wire input; never a laundering path.

## Zero dependencies

Stdlib only, by contract — stdlib-only consumers (gate batteries, CLIs, nix-hermetic builds) adopt it without a dependency discussion.

## Consumers

Adopted:

- `project-meta` doctor CI gate — the four gate categories roll up via `WorstOf` and the survey prints the typed verdict line (migrated 2026-11-21, `1e35099`; the library's first consumer, which per ADR-0001 makes it shipped)
- BuildFlow doctor survey — per-repo diagnostic statuses map onto PASS/WARN/FAIL/SKIP and the gate reads `WorstOf`; exit-code semantics identical (2026-11-21, `9a6af21ab`)

Evaluated and honestly closed (2026-11-21):

- monitor365 collector health — pure-Rust workspace (no Go side; its UI already has its own worst-combinator), no adoption possible

Note: docs-gates CANNOT adopt (stdlib-only root-module contract — zero requires, drift-checked in CI); its battery keeps its own three-word verdict vocabulary by design.
