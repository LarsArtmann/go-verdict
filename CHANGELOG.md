# Changelog

All notable changes. Dates are 2026 unless stated.

## [Unreleased]

### Added

- BuildFlow pre-commit hook gating every commit (block and auto-fix paths verified).
- GitHub Actions quality workflow (stdlib-only: build, vet, race tests on pushes and PRs; BuildFlow stays local-only as private tooling).
- Runnable godoc examples (`ExampleVerdict`, `ExampleWorstOf`, `ExampleParse`).
- FEATURES vocabulary-gate spike: statuses ⊆ own-legend check prototyped and measured; adopted upstream as Artmann-Technologies/docs-gates#11.

### Changed

- `weight()`'s roll-up ranks are named constants (`rankFail` … `rankSkip`) instead of bare numbers; the `.golangci.yml` `mnd` suppression for `verdict.go` is retired. No API change.
- `dprint.json` is no longer committed: BuildFlow generates it per host with local plugin paths (via `BUILDFLOW_DPRINT_PLUGINS`) where a vendored bundle exists, falling back to the remote first-use fetch elsewhere. Cold offline formatting runs are network-independent on bundled hosts.
- CONTRIBUTING.md now names the checks a contributor can actually run (build, vet, race tests — the same ones CI runs) instead of hand-run golangci-lint.

### Fixed

- Nothing yet.

## [v0.1.1] - 2026-11-22

### Added

- Documentation set built to fleet standard (docs-health audit, 2026-11-22): AGENTS.md, TODO_LIST.md, FEATURES.md, ROADMAP.md, and docs/DOMAIN_LANGUAGE.md; README installation section.
- `.golangci.yml` (BuildFlow fleet-default configuration; two curated suppressions documented in AGENTS.md).

### Fixed

- Consumer-record correction: project-meta doctor is the actual first consumer (migrated 2026-11-21); the planned first consumer docs-gates was contract-blocked and never adopted. Corrected in README and STACK.md (the v0.1.0 note below reflects the plan as it was believed at release time).

No code or API changes in this release; the library surface is untouched.

## [v0.1.0] - 2026-11-21

### Added

- The five-state `Verdict` type (PASS/FAIL/WARN/SKIP/UNKNOWN), `WorstOf` roll-up (Fail > Warn > Unknown > Pass > Skip; skips never vote; empty rolls to UNKNOWN), case-insensitive `Parse` that refuses near-miss laundering, `Icon` render grammar, and `ExitCode` (only FAIL exits 1). Stdlib-only. First consumer: docs-gates battery matrix.
