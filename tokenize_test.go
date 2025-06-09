package argv

import (
	"testing"

	"github.com/illbjorn/vsx/internal/zest"
)

func TestTokenize(t *testing.T) {
	z := zest.New(t)

	testTokenize(
		z,
		"empty",
		"",
		[]string{},
		false,
	)
	testTokenize(
		z,
		"bool flag",
		"-f",
		[]string{"-f"},
		false,
	)
	testTokenize(
		z,
		"string flag",
		"-f value",
		[]string{"-f", "value"},
		false,
	)
	testTokenize(
		z,
		"string flag with pos args",
		"-f value ok",
		[]string{"-f", "value", "ok"},
		false,
	)
	testTokenize(
		z,
		"non-terminated string literal",
		"-f 'value",
		[]string{},
		true,
	)
	testTokenize(
		z,
		"terminated string literal",
		"-f 'value 2'",
		[]string{"-f", "value 2"},
		false,
	)
}

func testTokenize(z zest.Zester, name, input string, want []string, wantErr bool) {
	z.Helper()

	z.With("input", input, "name", name)

	got, err := Tokenize(input)
	if wantErr {
		z.Assert(err != nil, "expected error")
	} else {
		z.Assert(err == nil, "expected no error, got [%s]", err)
	}

	z.Assert(
		len(got) == len(want),
		"expected [%d] values in Tokenize output, got [%d]",
		len(want), len(got),
	)

	for i := range min(len(got), len(want)) {
		z.Assert(
			got[i] == want[i],
			"expected [%s], got [%s] at position [%d] in tokenized output",
			want[i], got[i], i,
		)
	}
}
