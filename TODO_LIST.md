# TODO List

Open short- and mid-term work for go-verdict. Items are actionable only;
open decisions live in ROADMAP.md "Open questions", long-term raw ideas in
ROADMAP.md. Harvested from docs/status/2026-11-22_16-24_agents-md-and-hygiene-gate-green.md.

## Open

- [ ] Align CONTRIBUTING.md with BuildFlow ownership: it tells contributors to
      hand-run `golangci-lint run ./...` (CONTRIBUTING.md:17), which BuildFlow
      owns now.
- [ ] Add a godoc `ExampleWorstOf` example function (godoc shows no runnable
      example today).
- [ ] `buildflow precommit install` — install the pre-commit quality hook
      (none exists).
- [ ] BuildFlow state-DB maintenance (host-level): VACUUM
      ~/.local/state/buildflow/buildflow.db (1.09 GB; cache 0.60 GB) — also the
      cause of SQLITE_BUSY contention observed in three runs.
- [ ] Cut the next release: fold CHANGELOG `[Unreleased]` into a version
      section and tag (go-release skill).
- [ ] FEATURES-vocabulary standing-gate spike: prototype a check that every
      FEATURES.md status comes from its own legend (the 2026-11-22 audit
      shipped a legend gap no gate caught; surfaced as plan task P12 in
      docs/planning/2026-11-22_pareto-execution-plan.md).

## Watch items (not tasks)

- Verify pkg.go.dev renders the package doc after the next tag.
- Re-verify the README consumer ledger claims at their source repos next
  quarter (project-meta `1e35099` and BuildFlow `9a6af21` both verified
  2026-11-22).
- Revisit trigger from STACK.md: a render/SARIF layer growing past ~200 LOC
  outgrows the stdlib-only foundation decision — re-derive the stack before
  building it.
