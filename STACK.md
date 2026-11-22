# STACK.md — go-verdict

Decision record per the stack-selection skill. Profile-driven; re-derive before extending.

## Profile

| Axis        | Answer                                                                                                                                                       |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Shape       | Pure Go library (no CLI, no service)                                                                                                                         |
| Domain      | Typed epistemic status for gates/dashboards/check batteries                                                                                                  |
| Persistence | None                                                                                                                                                         |
| Interfaces  | Go API only (JSON-marshalable string type)                                                                                                                   |
| Constraints | **Zero dependencies, stdlib only — hard contract** (stdlib-only consumers like docs-gates must adopt it without dep discussions); nix-hermetic by triviality |

## Chosen foundation

| Concern    | Choice | Why this, not that                                  | Owning skill                                | Revisit trigger                               |
| ---------- | ------ | --------------------------------------------------- | ------------------------------------------- | --------------------------------------------- |
| Everything | stdlib | A five-state enum with a roll-up needs nothing else | how-to-golang (policy: smallest foundation) | a render/sarif layer that grows past ~200 LOC |

## Deliberately excluded

| Candidate               | Reason                                                                                                                                                                                                         |
| ----------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| host in go-health       | go-health is the K8s probe SDK with runtime deps (samber/do, go-branded-id); parking a stdlib-only concept there blocks its only guaranteed first consumer (docs-gates is stdlib-only by root-module contract) |
| host in go-output       | go-output owns formatting/output channels, not epistemics; hosting would couple state semantics to render plumbing                                                                                             |
| third-party status libs | no policy-approved generic candidate; the concept is 130 LOC                                                                                                                                                   |

## Wiring order

1. `meta set --project go-verdict` (registered 2026-11-21)
2. First consumer: project-meta doctor CI gate (migrated 2026-11-21, `1e35099`) — the planned first consumer docs-gates turned out contract-blocked (its own stdlib-only root-module forbids the import); per ADR-0001 the project-meta migration is what makes the library shipped
