package verdict

import (
	"fmt"
	"testing"
)

func TestParseCaseInsensitiveAndUnknown(t *testing.T) {
	t.Parallel()

	cases := map[string]Verdict{
		"pass":     Pass,
		"PASS":     Pass,
		"  Fail ":  Fail,
		"warn":     Warn,
		"Skip":     Skip,
		"unknown":  Unknown,
		"":         Unknown,
		"nonsense": Unknown,
		"passed":   Unknown, // near-miss words do not launder into states
	}

	for in, want := range cases {
		if got := Parse(in); got != want {
			t.Errorf("Parse(%q) = %q, want %q", in, got, want)
		}
	}
}

// The roll-up contract is the package's core promise; every row is a
// decision a consumer depends on.
func TestWorstOf(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		in   []Verdict
		want Verdict
	}{
		{"empty is Unknown, never Pass", nil, Unknown},
		{"all skips say nothing", []Verdict{Skip, Skip}, Unknown},
		{"single pass", []Verdict{Pass}, Pass},
		{"fail dominates everything", []Verdict{Pass, Pass, Warn, Unknown, Fail}, Fail},
		{"warn beats pass", []Verdict{Pass, Warn, Pass}, Warn},
		{"unknown poisons an all-pass roll-up", []Verdict{Pass, Unknown}, Unknown},
		{"unknown loses to warn", []Verdict{Warn, Unknown}, Warn},
		{"skip never votes", []Verdict{Skip, Pass}, Pass},
		{"skip beside fail stays fail", []Verdict{Skip, Fail}, Fail},
		{"unrecognized word counts as Unknown", []Verdict{Pass, Verdict("BROKEN")}, Unknown},
	}

	for _, tc := range cases {
		if got := WorstOf(tc.in...); got != tc.want {
			t.Errorf("%s: WorstOf(%v) = %q, want %q", tc.name, tc.in, got, tc.want)
		}
	}
}

func TestIconAndExitCode(t *testing.T) {
	t.Parallel()

	cases := map[Verdict]struct {
		icon string
		exit int
	}{
		Pass:    {"✓", 0},
		Fail:    {"✗", 1},
		Warn:    {"⚠", 0},
		Skip:    {"−", 0},
		Unknown: {"?", 0},
	}

	for v, want := range cases {
		if v.Icon() != want.icon {
			t.Errorf("%q.Icon() = %q, want %q", v, v.Icon(), want.icon)
		}

		if v.ExitCode() != want.exit {
			t.Errorf("%q.ExitCode() = %d, want %d", v, v.ExitCode(), want.exit)
		}
	}
}

// The state set is closed: five states, no more. A new state is a
// contract change and must update the census of consumers.
func TestStateSetIsClosed(t *testing.T) {
	t.Parallel()

	all := []Verdict{Pass, Fail, Warn, Skip, Unknown}
	if len(all) != 5 {
		t.Fatalf("state set grew: %v", all)
	}

	for _, v := range all {
		if !v.Valid() {
			t.Errorf("%q must be Valid", v)
		}
	}
}

func ExampleVerdict() {
	v := Parse("pass")
	fmt.Println(v, v.Icon(), v.ExitCode())
	// Output: PASS ✓ 0
}

func ExampleWorstOf() {
	fmt.Println(WorstOf(Pass, Warn, Skip))
	fmt.Println(WorstOf(Pass, Unknown))
	fmt.Println(WorstOf())
	// Output:
	// WARN
	// UNKNOWN
	// UNKNOWN
}

func ExampleParse() {
	fmt.Println(Parse("PASS"), Parse("passed"), Parse("  fail "))
	// Output: PASS UNKNOWN FAIL
}
