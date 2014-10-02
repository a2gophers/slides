// +build OMIT
package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	toServe := filepath.Join(
		os.Getenv("GOPATH"),
		"src",
		"github.com",
		"a2gophers",
	)

	// START OMIT
	m := martini.Classic()
	m.Get("/user/:id", func(w http.ResponseWriter, params martini.Params) {
		w.WriteHeader(http.StatusOK)
		fmt.Println("retrieving user with id:", params["id"])
	})

	http.HandleFunc("/ok", ok)
	http.Handle("/", http.FileServer(http.Dir(toServe)))

	http.ListenAndServe(":1810", m)
	// STOP OMIT
}

func ok(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
