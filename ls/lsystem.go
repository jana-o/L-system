//Package ls implements a system for generating drawings
package ls

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/index.html"))
}

//Lsystem is a parallel rewriting system
type Lsystem struct {
	Axiom               string
	RuleLeft, RuleRight string
	Result              []string
}

//NewLsystem creates instance of Lsystem
func NewLsystem(Axiom, RuleLeft, RuleRight string, Result []string) *Lsystem {
	l := &Lsystem{}
	l.Axiom = Axiom
	l.RuleLeft = RuleLeft
	l.RuleRight = RuleRight
	l.Result = Result

	return l
}

//Generate applies L-system rule n times
func (l *Lsystem) Generate() []string {
	s := l.Axiom
	n := len(l.Result) + 1

	for i := 0; i < n; i++ {
		s = l.applyRule(s)
	}

	l.Result = append(l.Result, s)

	fmt.Println("generate", l.Result[len(l.Result)-1], n)
	return l.Result
}

func (l *Lsystem) applyRule(r string) string {
	return strings.Replace(r, l.RuleLeft, l.RuleRight, -1)
}

func (l *Lsystem) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", l)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
