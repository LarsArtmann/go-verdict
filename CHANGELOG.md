# Changelog

All notable changes. Dates are 2026 unless stated.

## [v0.1.0] - 2026-11-21

### Added

- The five-state `Verdict` type (PASS/FAIL/WARN/SKIP/UNKNOWN), `WorstOf` roll-up (Fail > Warn > Unknown > Pass > Skip; skips never vote; empty rolls to UNKNOWN), case-insensitive `Parse` that refuses near-miss laundering, `Icon` render grammar, and `ExitCode` (only FAIL exits 1). Stdlib-only. First consumer: docs-gates battery matrix.
