// +build OMIT
package main

import (
	"github.com/codegangsta/martini"
	"net/http"
)

func main() {

	// START OMIT
	m := martini.New()

	// Specify what middleware to use
	m.Use(martini.Logger())
	m.Use(martini.Recovery())

	r := martini.NewRouter()
	r.Get("/ok", ok)
	r.Get("/panic", panicFunc)
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	http.ListenAndServe(":1810", m)
	// STOP OMIT
}

func panicFunc() {
	panic("we don't actually satisfy handler!!!")
}

func ok(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
