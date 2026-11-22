# TODO List

Open short- and mid-term work for go-verdict. Items are actionable only;
open decisions live in ROADMAP.md "Open questions", long-term raw ideas in
ROADMAP.md. Harvested from docs/status/2026-11-22_16-24_agents-md-and-hygiene-gate-green.md.

## Open

None. Open decisions live in ROADMAP.md (none open), long-term ideas in
ROADMAP.md themes.

## Watch items (not tasks)

- Re-verify the README consumer ledger claims at their source repos by
  2027-02-22 (project-meta `1e35099` → cmd/meta/cmd_doctor.go and BuildFlow
  `9a6af21` → internal/cli/doctor_survey_verdict.go, both verified
  2026-11-22).
- When docs/status reaches three reports, add an index page for it (two
  reports today).
- Revisit trigger from STACK.md: a render/SARIF layer growing past ~200 LOC
  outgrows the stdlib-only foundation decision — re-derive the stack before
  building it.
