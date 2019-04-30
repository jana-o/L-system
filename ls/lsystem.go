//Package ls implements a system for generating drawings
package ls

import (
	"fmt"
	"strings"
)

//Lsystem is
type Lsystem struct {
	axiom               string
	ruleLeft, ruleRight string
}

//NewLsystem creates instance of Lsystem
func NewLsystem(axiom, ruleLeft, ruleRight string) *Lsystem {
	l := &Lsystem{}
	l.axiom = axiom
	l.ruleLeft = ruleLeft
	l.ruleRight = ruleRight

	return l
}

//Generate applies L-system rule n times
func (l *Lsystem) Generate(n int) string {
	s := l.axiom
	for i := 0; i < n; i++ {
		s = l.applyRule(s)
	}
	fmt.Printf("%T\n%v\n", s, s)
	return s
}

func (l *Lsystem) applyRule(r string) string {

	return strings.Replace(r, l.ruleLeft, l.ruleRight, -1)

}
