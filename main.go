package main

import (
	"code/lsystem-v6/ls"
	"code/lsystem-v6/turtlego"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	l := ls.NewLsystem("F", "F", "FF+[+F-F-F]-[-F+F+F]", []string{})

	http.Handle("/", l)
	http.Handle("/create", create())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		l := ls.NewLsystem("F", "F", "FF+[+F-F-F]-[-F+F+F]", []string{})

		err := req.ParseForm()
		if err != nil {
			panic(err)
		}
		n := req.Form.Get("n")
		i, err := strconv.Atoi(n)

		l.Generate(i)
		turtlego.ToPNG(l, fmt.Sprintf("./public/pic/l-system%d.png", i))
		l.ServeHTTP(w, req)
	})
}
