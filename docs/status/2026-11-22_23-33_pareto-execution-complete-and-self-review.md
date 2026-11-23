# Status Report — Pareto plan executed (P1–P13), v0.1.1 shipped — 2026-11-22 23:33

**Scope line.** Verified here: every claim below was checked live this session —
`buildflow` (exit 0) and `docs-gates run .` (5/5 PASS) at every commit tip;
`go test ./... -race -count=1` green at `945e6be` and `324f200`; release
verified via `go get` in a clean module, proxy `.info` (hash `9545aed…`), the
pkg.go.dev versions tab (v0.1.1 listed), and `gh release list`; CI verified
green on runs `45054250429`, `45054375749`, and `45055822387` (the final tip);
consumer ledger verified by grep at both source repos. NOT verified here: the
auto-commit daemon passing through the new pre-commit hook (no daemon commit
occurred this session); the recurring buildflow "Rebuild and reinstall" doctor
warning's root cause (observed, not diagnosed); the `timings --info` footer
line "type: NOT IMPLEMENTED" (unexplained); gitleaks/codespell (on-demand
tools, not run); pkg.go.dev's docs body (withheld by license policy — only the
version listing was verified); test coverage percentage (not measured).

Session arc: executed docs/planning/2026-11-22_18-46_pareto-execution-plan.md
tasks P2–P13 (P1 landed pre-plan). Owner verdicts collected mid-session via
question tool: refactor weight() / local dprint plugins / Actions CI workflow.

## a) Fully done

- **P2 — Release v0.1.1** (`9545aed`, tag `v0.1.1`): CHANGELOG folded,
  full gates green pre-tag, annotated tag pushed, GitHub Pre-release created
  (2026-11-22T18:08Z), `go get github.com/larsartmann/go-verdict@v0.1.1`
  resolves/builds/runs in a clean temp module (roll-up smoke test printed
  `FAIL` correctly), proxy.golang.org indexed (hash match), pkg.go.dev lists
  v0.1.1.
- **P3 — Owner decisions** (`4e7efb9`): three verdicts recorded, all three
  ROADMAP Open questions closed inline with outcomes.
- **P4 — Contributor trust** (`324f200`): CONTRIBUTING.md now names the checks
  a contributor can actually run; `ExampleVerdict`/`ExampleWorstOf`/
  `ExampleParse` godoc examples added (output-verified by `go test`).
- **P5 — Pre-commit hook** (`a9dc57c`): `buildflow precommit install`; block
  path verified (compile error aborted a scratch commit), auto-fix path
  verified (misaligned table committed as dprint-formatted), real commits pass
  through it; documented in AGENTS.md with daemon interplay.
- **P6 — Suppressions verdict executed** (`945e6be`): `weight()` uses named
  rank constants (`rankFail`…`rankSkip`); `mnd` suppression for `verdict.go`
  dropped with zero golangci findings on that file.
- **P7 — Local dprint plugins** (`56d61fa`): `dprint.json` untracked and
  gitignored, generated per host with local plugin paths (host bundle already
  exported by home-manager — discovered after a wasted duplicate provision);
  cold-cache dead-proxy run green; AGENTS.md documents the pattern.
- **P8 — CI workflow** (`4e7efb9`, amended): `.github/workflows/quality.yml`
  stdlib-only (build/vet/race), SHA-pinned by the auto-fixer after the gate
  rejected my tag pins; dependabot gained the github-actions ecosystem; three
  green runs verified including the final tip `69456e0`.
- **P9 — Host maintenance** (no commit — host-level): 620 MB result cache
  purged; state DB diagnosed healthy via `timings --info` (retention working,
  oldest record 1 month, 0% fragmentation) so VACUUM was correctly skipped;
  post-purge run green, no SQLITE_BUSY.
- **P10 — Quarterly ritual** (`d5e3958`): consumer ledger re-verified at both
  sources (project-meta `1e35099` → cmd/meta/cmd_doctor.go; BuildFlow
  `9a6af21` → internal/cli/doctor_survey_verdict.go); next check 2027-02-22 in
  TODO watch items.
- **P11 — Watch triggers** (`d5e3958`): census re-run at five adopters
  (ROADMAP); docs/status index at three reports (TODO watch).
- **P12 — Vocabulary-gate spike** (`69456e0`): prototype validated (detects a
  synthetic legend gap; false-positive class measured = non-status ALL-CAPS
  tokens in other columns), adopted upstream as
  Artmann-Technologies/docs-gates#11 (verified landed via issue list).
- **P13 — Demand gate** (`69456e0`): render/SARIF demand check recorded in
  ROADMAP — two adopters, neither duplicates a render layer; stays gated.
- **Report hygiene**: 27 report items annotated inline across both status
  reports (emit-keys → verify → write; both freshness and audit-all FRESH).
- Final state: master `69456e0` pushed, tree clean, TODO_LIST open section
  empty, all gates green, CI green on the final tip.

## b) Partially done

- **Daemon × pre-commit-hook coexistence** (P5.4): the hook gates every commit
  and real commits pass, but NO auto-commit-daemon commit occurred this
  session, so the documented expectation (blocked daemon commits leave changes
  in the tree; hook repairs can surface inside daemon commits) is
  documentation, not observation. Effort to finish: S (wait for a daemon
  commit, verify with `git log --stat`).
- **CI-on-release verification**: v0.1.1 predates the workflow, so no release
  tag has ever been CI-verified (go-release skill Phase 4.4 wants the tagged
  commit green). Next release must check `gh run list` before tagging.
  Effort: S (part of the next release ritual).
- **The 18-12 report's §g questions** (the three owner questions) were
  answered and closed in ROADMAP, but the §g section itself carries no inline
  resolution pointer (freshness only audits §f). Honesty gap, not a gate
  failure. Effort: S (one routed marker per question).

## c) Not started

- **docs-gates#11 implementation**: the vocabulary gate is filed upstream with
  a validated design; no code written here or there (deliberate — fleet tool,
  belongs in docs-gates).
- **Plan-file inline verdicts**: docs/planning/2026-11-22_18-46_pareto-
  execution-plan.md still shows P2–P13 as future work; completion is recorded
  in commit messages and this report, not annotated in the plan itself.
- **v0.1.2**: `[Unreleased]` accumulates (hook, CI, examples, CONTRIBUTING,
  dprint, ranks) but no cut decision made — see §g question 1.
- **Render/SARIF exploration** (P13 work itself): not started by design —
  demand-gated, no consumer signal.

## d) Totally fucked up (all caught and recovered; nothing shipped broken)

1. **Committed on a red gate (P8).** I chained `buildflow; echo; git add &&
   git commit` — the `;` let the commit run although buildflow exited 69
   (action tag pins = error severity). Found it by reading the exit code
   afterward; fixed via the auto-fixer and amended before push. Root cause:
   habit-form command chaining; the fix is structural — a commit never runs
   unless the gate command before it succeeded (`&&`, not `;`).
2. **Deleted TODO rows before annotating the source reports** — freshness and
   audit-all went red (BF=69, DG=1, five unharvested §f items). I knew the
   canon (annotate on completion) and even loaded the docs-health skill later,
   but edited the living docs first. Worse: `docs-gates freshness --pre-flight`
   exists exactly for this and I did not run it. Recovered with a 27-item
   annotate sweep; both reports FRESH.
3. **Provisioned a duplicate dprint bundle before checking the ambient
   environment.** I nix-built and copied 3.5 MB of wasm to
   ~/.local/share/dprint-plugins, attempted a .bashrc append (denied —
   home-manager store symlink, which global AGENTS.md documents), and only
   then discovered home-manager ALREADY exports BUILDFLOW_DPRINT_PLUGINS at
   the exact path the doctor names. Check-ambient-before-provisioning would
   have saved the whole detour. Duplicate cleaned up.
4. **Question tool schema fumbled three times** (missing `type`, then missing
   `description` on every question) before a valid batch. Two wasted round
   trips; the schema needs `type` + `description` on every question.
5. **`rg -rn` flag trap hit a THIRD time on this host** — the hook that
   blocked it says so. `rg -n` is already recursive; `-r` is replace-display.
6. **Guessed CI action majors** (`checkout@v5`, `setup-go@v6` — reality:
   v7.0.1/v7.1.0). The github-actions-pinning gate caught it; guessing
   versions instead of resolving them was lazy.
7. **Minor sloppiness**: a bad glob (`./example...`) in the release smoke
   test caused a confusing exit 1 (rerun clean); a `printf` with parens
   broke mvdan/sh quoting (switched to the write tool); a malformed
   `sed '^$d'` in the prototype (output unaffected because the comm sets
   were empty).

## e) What we should improve

- **Gate-then-commit must be one `&&` chain.** The P8 incident is the second
  time this session a `;` separated verification from mutation. Make it a
  reflex: no commit command runs unless the gate before it exits 0.
- **Run `docs-gates freshness --pre-flight` before every living-doc edit.**
  It exists precisely to prevent the (d)(2) failure and I skipped it twice.
- **Annotate reports in the same step as completing the work.** Completion is
  two writes (living doc + historical report), not one; doing only the first
  always goes red.
- **Check ambient host state before provisioning anything** (env vars,
  bundles, dotfiles — which are home-manager store symlinks here, read-only
  by design).
- **Close the loop on observed-but-unexplained output.** Two sightings were
  noted and skipped: the recurring buildflow "Rebuild and reinstall" doctor
  warning (appears in every full run's log) and `timings --info` printing
  "type: NOT IMPLEMENTED". Both are upstream (BuildFlow) observations; each
  deserves a diagnose → fix/file/document decision, not a shrug.
- **Resolve versions, never guess them** (actions SHAs, tool versions) — the
  gates catch it, but guessing wastes a cycle every time.

## f) Next things (ranked; feeds HARVEST)

1. Decide v0.1.2 timing (owner question — see §g) and, when cutting: verify CI
   green on the tagged commit. Impact High, Effort S, Release.
2. Observe one real auto-commit-daemon commit pass the pre-commit hook;
   verify with `git log --stat`. Impact Medium, Effort S, Quality.
3. Diagnose the recurring buildflow "Rebuild and reinstall" doctor warning;
   fix via the documented reinstall path or file upstream. Impact Medium,
   Effort S, Tooling.
4. Investigate `buildflow timings --info` footer "type: NOT IMPLEMENTED"
   (upstream BuildFlow; verify at source, then fix or file). Impact Low,
   Effort S, Tooling.
5. Annotate the 18-12 report's §g questions with routed markers pointing at
   the ROADMAP closure. Impact Low, Effort S, Documentation.
6. Annotate the plan file's P1–P13 rows with done verdicts (planning docs are
   historical; annotate, never rewrite). Impact Low, Effort S, Documentation.
7. Implement docs-gates#11 (vocabulary gate) upstream: column-aware status
   extraction on the existing rows scanner; legend ⊆ check. Impact Medium,
   Effort M, Tooling.
8. Add `concurrency: cancel-in-progress` to quality.yml so superseded pushes
   don't queue 8-minute runs. Impact Low, Effort S, CI.
9. Review the first dependabot github-actions PR: confirm SHA pins update
   cleanly under the pinning gate. Impact Low, Effort S, CI.
10. Add `ExampleIcon` and `ExampleExitCode` godoc examples for full surface
    coverage (only if they teach something the three existing examples do
    not). Impact Low, Effort S, Documentation.
11. README: one line noting pkg.go.dev withholds rendered docs under the
    proprietary license (godoc examples exist; read them in-repo). Impact
    Low, Effort S, Documentation.
12. Decide on a version badge in README (none exists; license badge only).
    Impact Low, Effort S, Documentation.
13. Consider PR/issue templates (.github/). Impact Low, Effort S, CI.
14. Consider measuring coverage once (`go test -cover`) and recording it in
    FEATURES evidence columns. Impact Low, Effort S, Quality.
15. When a third adopter lands: update README consumer ledger and the census
    trigger counter (ROADMAP says re-run census at five). Impact Medium,
    Effort S, Documentation.
16. Keep the quarterly re-verification appointment 2027-02-22 (TODO watch
    item; the nextd queue could carry it as a reminder if the owner wants).
    Impact Medium, Effort S, Process.
17. If a render/SARIF request ever arrives from a real consumer: re-derive
    the stack first (STACK.md guard rail, >200 LOC trigger). Impact
    Conditional, Effort L, Feature.
18. Add a `docs/status/` index page when the third report lands (watch item).
    Impact Low, Effort S, Documentation.
19. Next go-release: run the skill's pre-push checks including `gh run list`
    on the tag commit (this release had no CI to check). Impact Medium,
    Effort S, Release.
20. Consider `GOTOOLCHAIN` pinning in CI only if a stdlib CVE chase ever
    forces it (go-release skill tip; do not pre-build). Impact Low, Effort
    S, CI.

(20 items, not 50: the honest backlog is short — the plan drained it. Padding
to 50 would manufacture work.)

## g) Questions I cannot answer myself

1. **v0.1.2 timing.** `[Unreleased]` holds only repo-infrastructure entries
   (hook, CI, examples, CONTRIBUTING, dprint config, ranks refactor — no API
   change). Cut a patch now so consumers pin current docs, or wait for a real
   library delta? I could argue either way; release cadence is owner taste.
2. **Host-level buildflow upgrade.** Every full run prints a doctor warning
   ("Rebuild and reinstall: nix build . && nix run .#reinstall … BuildFlow#81").
   Rebuilding/reinstalling the shared buildflow binary affects every fleet
   repo's gate — do you want me to run it (and when: now, or after this
   session's work is merged), or leave it to your host ritual?
3. **docs/drafts/ retention.** The github-voice skill mandates keeping filing
   drafts in-repo (never /tmp), so docs-gates#11's draft is committed at
   docs/drafts/2026-11-22_docs-gates-vocabulary-gate-issue.md. This repo's
   ethos is minimal — keep the drafts dir as living provenance, or trash
   drafts once their issues are verified filed?

---

Provenance: written by Crush (glm-5.3) at the owner's request immediately
after session close; all hashes/run-IDs re-read from this session's command
outputs, not from memory of prior sessions.
