# Draft: docs-gates vocabulary-gate issue

## Title

gates gap: nothing checks that FEATURES.md statuses come from the file's own legend

## Body

## Problem

The battery verifies doc shape (freshness, annotate, routing, cadence, rows) but not vocabulary: a FEATURES.md row can carry a status token the file's own legend never defines, and every gate stays green. That exact gap shipped in the wild — the 2026-11-22 go-verdict docs audit published a FEATURES.md whose "Deliberately absent" table used `OUT_OF_SCOPE` before any legend defined it (caught in self-review, fixed in the same session; no gate would have flagged it).

## Evidence

Prototype against a live FEATURES.md (go-verdict, 2026-11-22):

- legend tokens extracted from the `## Status legend` section (table + absent-why prose note): `BROKEN`, `DISABLED`, `FULLY_FUNCTIONAL`, `OUT_OF_SCOPE`, `PARTIALLY_FUNCTIONAL`, `PLANNED`
- all backticked ALL-CAPS tokens used anywhere in the file: identical set — clean
- synthetic mutation of one row's status to `IN_PROGRESS`: flagged as the only violation
- false-positive class measured: whole-line scanning also flags legitimate ALL-CAPS backticked tokens in non-status columns (e.g. `JSON` in a Notes cell) — so the gate needs status-column-aware extraction, not a whole-line grep

## Proposal

A vocabulary gate (say `docs-vocabulary`, joining the battery): every table's status column must only contain tokens the file's own legend defines; legend-unused is informational, never a failure.

|                        | today | with the gate |
| ---------------------- | ----- | ------------- |
| undefined status used  | green | FAIL          |
| status-cell extraction | none  | column-aware  |
| legend drift           | green | info note     |

The seam: the rows/bare-rows scanner already walks tables cell-aware — the status-column extraction can reuse it rather than a second parser. Where the legend lives (table, prose extension like the absent-why note) is the design question worth settling: the go-verdict file carries one status (`OUT_OF_SCOPE`) defined only in prose under the legend heading.

- [ ] vocabulary check: status tokens ⊆ legend set, FAIL on violation
- [ ] legend extraction covers prose extensions, not just the legend table
- [ ] unused-legend-entry report stays informational

Origin: spike task P12 of the go-verdict 2026-11-22 execution plan (docs/planning/2026-11-22_18-46_pareto-execution-plan.md), disposition "adopt in a gate" — fleet-wide value belongs here, not as project-local tooling.

💘 Generated with Crush
