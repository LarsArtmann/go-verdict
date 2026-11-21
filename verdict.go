// Package verdict is the fleet's typed epistemic status: one small set
// of states, one worst-of roll-up, one render grammar.
//
// Born from the 2026-11-21 domain-language census: ~14 dashboards and
// gate batteries hand-roll the same concept — a check ran, and its
// outcome is one of a few honest states — each with its own strings,
// its own roll-up ordering, and its own glyphs. This package is that
// concept, extracted once.
//
// The states, and what each one PROMISES:
//
//   - Pass:     ran, and the claim it tests holds.
//   - Fail:     ran, and the claim is false. The only failing state.
//   - Warn:     ran, the claim holds but carries debt the reader must
//     see (a warn-mode gate finding, a near-threshold drift). Never
//     fails a roll-up by itself.
//   - Skip:     deliberately not evaluated here (exempt, filtered, not
//     applicable). Skips do not vote in a roll-up.
//   - Unknown:  cannot say — never ran, timed out, or evidence was
//     inconclusive. The honest default; never silently coerced to Pass.
//
// Roll-up (WorstOf): Fail > Warn > Unknown > Pass > Skip. A Fail in any
// input fails the roll-up; a lone Warn warns; Unknown poisons a would-be
// all-Pass roll-up to Unknown (unverified is not verified); Skips never
// vote — a roll-up of only Skips is Unknown, because nothing was said.
//
// Zero dependencies, stdlib only — the same contract that lets
// stdlib-only consumers (gate batteries, CLIs) adopt it.
package verdict

import "strings"

// Verdict is one check's epistemic status. The zero value is the empty
// string; Parse maps it (and any unknown word) to Unknown — an empty
// verdict is an unmeasured one, never a passing one.
type Verdict string

// The complete state set. String forms are upper-case words, stable
// wire format (JSON, SARIF, TSV) across every consumer.
const (
	Pass    Verdict = "PASS"
	Fail    Verdict = "FAIL"
	Warn    Verdict = "WARN"
	Skip    Verdict = "SKIP"
	Unknown Verdict = "UNKNOWN"
)

// weight is the roll-up rank. Higher outranks; Skip does not vote
// (represented as the lowest weight and skipped in WorstOf).
func (v Verdict) weight() int {
	switch v {
	case Fail:
		return 5
	case Warn:
		return 4
	case Unknown:
		return 3
	case Pass:
		return 2
	case Skip:
		return 1
	}

	return 3 // unrecognized verdicts carry Unknown's weight: unmeasured
}

// Valid reports whether v is one of the five defined states.
func (v Verdict) Valid() bool {
	switch v {
	case Pass, Fail, Warn, Skip, Unknown:
		return true
	}

	return false
}

// Parse maps any string to a Verdict, case-insensitively; anything
// unrecognized (including the empty string) becomes Unknown. Use for
// wire input; never use to launder an unknown word into a pass.
func Parse(s string) Verdict {
	v := Verdict(strings.ToUpper(strings.TrimSpace(s)))
	if v.Valid() {
		return v
	}

	return Unknown
}

// WorstOf rolls up verdicts: Fail > Warn > Unknown > Pass > Skip.
// Skips never vote: they are dropped before the maximum. An empty list
// (or all-Skip list) rolls up to Unknown — nothing was said, and an
// honest roll-up does not invent a Pass.
func WorstOf(vs ...Verdict) Verdict {
	worst := Skip

	for _, v := range vs {
		if v == Skip {
			continue
		}

		if !v.Valid() {
			v = Unknown // unrecognized inputs roll up AS Unknown, never as their raw word
		}

		if v.weight() > worst.weight() {
			worst = v
		}
	}

	if worst == Skip {
		return Unknown
	}

	return worst
}

// Icon is the render grammar: one glyph per state, terminal-stable,
// matching the fleet's dashboards (doctor's ✓/⚠ convention extended).
func (v Verdict) Icon() string {
	switch v {
	case Pass:
		return "✓"
	case Fail:
		return "✗"
	case Warn:
		return "⚠"
	case Skip:
		return "−"
	case Unknown:
		return "?"
	}

	return "?"
}

// ExitCode maps a verdict to the process exit contract the fleet's
// CLIs share: only Fail is an error exit (1); everything else exits 0.
// Consumers that want warn-mode gates to exit non-zero compose that
// policy themselves — the type carries epistemics, not policy.
func (v Verdict) ExitCode() int {
	if v == Fail {
		return 1
	}

	return 0
}
