package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/index.html"))
}

type Lsystem struct {
	Axiom               string
	RuleLeft, RuleRight string
	Result              []string
}

func (l *Lsystem) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", l)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

//Generate applies L-system rule n times
func (l *Lsystem) Generate() []string {
	s := l.Axiom
	n := len(l.Result) + 1

	for i := 0; i < n; i++ {
		s = l.applyRule(s)
	}

	l.Result = append(l.Result, s)

	fmt.Println("generate", s, l.Result, n)
	return l.Result
}

func (l *Lsystem) applyRule(r string) string {
	return strings.Replace(r, l.RuleLeft, l.RuleRight, -1)
}

func main() {
	l := &Lsystem{
		"F",
		"F",
		"F+F-",
		[]string{},
	}
	// l := ls.NewLsystem("F", "F", "F+F-")
	// turtlego.ToPNG(l, "./public/pic/l-system.png")

	http.Handle("/", index(l))
	http.Handle("/create", create(l, l))

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		h.ServeHTTP(w, req)
	})
}

func create(h http.Handler, l *Lsystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		l.Generate()
		// fmt.Println("create", l.Result)
		h.ServeHTTP(w, req)
	})
}
