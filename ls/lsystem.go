//Package ls implements a system for generating drawings
package ls

import (
	"fmt"
	"strings"
)

//Lsystem is
type Lsystem struct {
	Axiom               string
	RuleLeft, RuleRight string
	// Result              string
}

//NewLsystem creates instance of Lsystem
func NewLsystem(Axiom, RuleLeft, RuleRight string) *Lsystem {
	l := &Lsystem{}
	l.Axiom = Axiom
	l.RuleLeft = RuleLeft
	l.RuleRight = RuleRight

	return l
}

//Generate applies L-system rule n times
func (l *Lsystem) Generate(n int) string {
	s := l.Axiom
	// r := l.Result
	for i := 0; i < n; i++ {
		s = l.applyRule(s)
	}
	fmt.Printf("%T\n%v\n", s, s)
	// s = r
	// fmt.Println("generate", &s, r)
	return s
}

func (l *Lsystem) applyRule(r string) string {
	return strings.Replace(r, l.RuleLeft, l.RuleRight, -1)
}
