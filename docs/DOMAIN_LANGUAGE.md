# Domain Language

> The ubiquitous language of this project. Every term used in code,
> conversations, and docs should be defined here unambiguously.
>
> Born from the 2026-11-21 fleet domain-language census: the same concept was
> hand-rolled ~14 times with private vocabularies. This glossary is the
> extracted shared one.

## Core terms

| Term                 | Definition                                                                                                       | Where used in code                           |
| -------------------- | ---------------------------------------------------------------------------------------------------------------- | -------------------------------------------- |
| Verdict              | One check's epistemic status: what the check can honestly claim about the world                                  | `Verdict` type, `verdict.go:36`              |
| Epistemic status     | What a check can say — ran and holds, ran and is false, ran with debt, deliberately not evaluated, or cannot say | Package doc, `verdict.go:1-29`               |
| Pass                 | Ran, and the claim it tests holds                                                                                | `Pass` constant, `verdict.go:41`             |
| Fail                 | Ran, and the claim is false — the only failing state                                                             | `Fail` constant, `verdict.go:42`             |
| Warn                 | Ran, holds, but carries debt the reader must see; never fails a roll-up alone                                    | `Warn` constant, `verdict.go:43`             |
| Skip                 | Deliberately not evaluated here (exempt, filtered, not applicable)                                               | `Skip` constant, `verdict.go:44`             |
| Unknown              | Cannot say — never ran, timed out, or inconclusive; the honest default, never silently coerced to Pass           | `Unknown` constant, `verdict.go:45`          |
| Roll-up              | Combining many verdicts into one via worst-wins ordering: Fail > Warn > Unknown > Pass > Skip                    | `WorstOf`, `verdict.go:93`                   |
| Voting               | Which states participate in a roll-up; Skips never vote — they are dropped before the maximum                    | `WorstOf` doc comment, `verdict.go:89-92`    |
| Honest roll-up       | A roll-up that does not invent a Pass: empty or all-Skip input is Unknown (nothing was said)                     | `WorstOf` doc comment, `verdict.go:90-92`    |
| Poisoning            | Unknown poisons an otherwise all-Pass roll-up — unverified is not verified                                       | Package doc, `verdict.go:22-25`              |
| Wire format          | The stable upper-case string forms (PASS/FAIL/WARN/SKIP/UNKNOWN) used across JSON, SARIF, TSV                    | Constants block doc, `verdict.go:38-39`      |
| Laundering           | Turning an unrecognized word into a state; `Parse` refuses it — near-misses ("passed") become Unknown            | `Parse` doc comment, `verdict.go:77-79`      |
| Render grammar       | One glyph per state, terminal-stable, shared across the fleet's dashboards                                       | `Icon`, `verdict.go:119`                     |
| Exit contract        | The process-level meaning of a verdict: only Fail is an error exit (1); everything else exits 0                  | `ExitCode`, `verdict.go:140`                 |
| Epistemics vs policy | The type reports what was measured; consequences (fail the build on Warn, retries, gating) belong to consumers   | `ExitCode` doc comment, `verdict.go:136-139` |
| Zero value           | The empty-string Verdict; an unmeasured verdict, mapped to Unknown by Parse — never a passing one                | `Verdict` doc comment, `verdict.go:33-35`    |

## Bounded contexts

None — single-context library; every term has exactly one meaning.

## Deprecated terms

None — the vocabulary was extracted once (2026-11-21 census); no retired terms
yet. The private per-dashboard vocabularies this package replaces (ok/err
booleans, three-word gate vocabularies) are deprecated in those consumers,
not here.
