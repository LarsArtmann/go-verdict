# Roadmap

> Long-term direction and raw ideas. Items here are NOT actionable tasks.
> When an idea is refined into bounded work, it moves to TODO_LIST.md.

## Themes

### 1. Fleet adoption

The library exists because ~14 dashboards and gate batteries hand-rolled the
same concept with divergent strings, orderings, and glyphs. Adoption is the
product: every migration deletes a private vocabulary.

Raw ideas:

- Migrate the remaining hand-rolled verdict vocabularies across the fleet
  (consumer ledger lives in README.md)
- Re-run the domain-language census as adoption grows, to catch duplicate
  glossaries and near-synonyms the migrations leave behind
- Quarterly re-verification of the README consumer ledger claims at their
  source repos

### 2. Render surface (demand-gated)

The type carries epistemics; rendering is a consumer concern today. A shared
render layer only earns existence once multiple consumers duplicate one.

Raw ideas:

- SARIF emission for CI-native consumers
- Table/summary renderers for doctor-style surveys
- Guard rail from STACK.md: a render/SARIF layer growing past ~200 LOC
  outgrows the stdlib-only foundation decision — re-derive the stack before
  building it, not after

## Non-goals

- **A sixth state:** the five-state set is closed; growth is a breaking
  contract change, not a feature.
- **Policy in the type:** warn-as-error exit behavior, deployment gating,
  retry logic — consumers compose those themselves.
- **Third-party dependencies:** stdlib-only is the adoption contract that lets
  stdlib-only consumers adopt without a dependency discussion.
- **A CLI or service:** pure library; wrapping CLIs live in their consumers.

## Open questions

None currently open. The three standing questions were decided on
2026-11-22 and executed: the golangci suppressions question was resolved by
naming `weight()`'s roll-up ranks (retiring the `mnd` suppression); dprint
plugins are provisioned locally (`BUILDFLOW_DPRINT_PLUGINS`, generated
`dprint.json`); CI is a stdlib-only GitHub Actions workflow (private
BuildFlow cannot run on public runners).
