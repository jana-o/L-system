package main

import (
	"code/lsystem-v6/ls"
	"code/lsystem-v6/turtlego"
	"log"
	"net/http"
)

func main() {

	l := ls.NewLsystem("F", "F", "FF+[+F-F-F]-[-F+F+F]", []string{})

	http.Handle("/", l)
	http.Handle("/create", create(l, l))
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func create(h http.Handler, l *ls.Lsystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		l.Generate()
		turtlego.ToPNG(l, "./public/pic/l-system.png")
		l.ServeHTTP(w, req)
	})
}

//FF+[+F-F-F]-[-F+F+F]
