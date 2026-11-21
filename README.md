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

| State     | Promise                                             | Fails a roll-up? |
| --------- | --------------------------------------------------- | ---------------- |
| `PASS`    | ran, and the claim it tests holds                   | no               |
| `FAIL`    | ran, and the claim is false — the only failing state | **yes**          |
| `WARN`    | ran, holds, but carries debt the reader must see    | no               |
| `SKIP`    | deliberately not evaluated here (exempt/filtered)   | no (never votes) |
| `UNKNOWN` | cannot say — never ran, timed out, inconclusive     | no               |

## Roll-up (`WorstOf`)

`Fail > Warn > Unknown > Pass > Skip`. Skips never vote; empty or all-Skip input rolls up to `UNKNOWN` (nothing was said — an honest roll-up does not invent a Pass). An `UNKNOWN` poisons an otherwise all-Pass roll-up: unverified is not verified.

## Parse

`Parse(" fail ")` → `FAIL`; any unrecognized word (including `""` and near-misses like `"passed"`) → `UNKNOWN`. For wire input; never a laundering path.

## Zero dependencies

Stdlib only, by contract — stdlib-only consumers (gate batteries, CLIs, nix-hermetic builds) adopt it without a dependency discussion.

## Consumers

- `docs-gates` battery matrix (first adopter, 2026-11-21)
