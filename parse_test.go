package argv

import (
	"testing"

	"github.com/illbjorn/vsx/internal/zest"
)

func TestParse(t *testing.T) {
	z := zest.New(t)

	testParse(
		z,
		"empty",
		[]string{},
		Command{},
		false,
	)
	testParse(
		z,
		"bool flag",
		[]string{"-f"},
		Command{Flags: map[string][]string{"f": {"true"}}},
		false,
	)
	testParse(
		z,
		"string flag",
		[]string{"-f", "value"},
		Command{Flags: map[string][]string{"f": {"value"}}},
		false,
	)
	testParse(
		z,
		"command with string flag",
		[]string{"-f", "value", "ok"},
		Command{Name: "ok", Flags: map[string][]string{"f": {"value"}}},
		false,
	)
	testParse(
		z,
		"command with string flag and pos args",
		[]string{"cmd", "-f", "value", "ok"},
		Command{Name: "cmd", Flags: map[string][]string{"f": {"value"}}, Args: []string{"ok"}},
		false,
	)
	testParse(
		z,
		"command with bool flag and pos args",
		[]string{"cmd", "ok", "-f"},
		Command{Name: "cmd", Flags: map[string][]string{"f": {"true"}}, Args: []string{"ok"}},
		false,
	)
}

func testParse(z zest.Zester, name string, input []string, want Command, wantErr bool) {
	z.Helper()
	z.With("input", input, "name", name)

	// Execute the parse
	got, err := Parse(input)
	if wantErr {
		z.Assert(err != nil, "expected error")
	} else {
		z.Assert(err == nil, "expected no error, got [%s]", err)
	}

	// Compare command name
	z.Assert(
		got.Name == want.Name,
		"expected command name [%s], got [%s]",
		want.Name, got.Name,
	)

	// Compare flags
	for name, values := range want.Flags {
		gotValues := got.Flag(name)
		// Confirm we got the expected number of flag values
		z.Assert(
			len(gotValues) == len(values),
			"expected [%d] values for flag [%s], got [%d]",
			len(values), name, len(gotValues),
		)
		// Confirm the flag values match expectations
		for i, value := range values {
			z.Assert(
				gotValues[i] == value,
				"expected flag value [%s] at index [%d], got [%s]",
				value, i, gotValues[i],
			)
		}
	}

	// Compare positional args
	z.Assert(
		len(got.Args) == len(want.Args),
		"expected [%d] args, got [%d]",
		len(want.Args), len(got.Args),
	)
	for i := range min(len(got.Args), len(want.Args)) {
		z.Assert(
			got.Args[i] == want.Args[i],
			"expected arg value [%s] at index [%d], got [%s]",
			want.Args[i], i, got.Args[i],
		)
	}
}
