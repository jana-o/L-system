package ls

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	l := Lsystem{}
	l.Axiom = "F"
	l.RuleLeft = "F"
	l.RuleRight = "F+F"
	y := l.Generate(2)
	if y != "F+F+F+F" {
		t.Error("expected F+F", "got", y)
	}
}

func TestApplyRule(t *testing.T) {
	l := Lsystem{}
	l.Axiom = "F"
	l.RuleLeft = "F"
	l.RuleRight = "FF"

	cases := []struct {
		input string
		want  string
	}{
		{"F", "FF"},
		{"[F+F]", "[FF+FF]"},
	}
	for _, c := range cases {
		got := l.applyRule(c.input)
		if got != c.want {
			t.Errorf("l.applyRule(%s) == %s, want %s", c.input, got, c.want)
		}
	}

}
