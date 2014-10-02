package main

import (
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	toServe := filepath.Join(
		os.Getenv("GOPATH"),
		"src",
		"github.com",
		"a2gophers",
	)

	http.HandleFunc("/ok", ok)
	http.HandleFunc("/random-num", randomNum)
	http.Handle("/", http.FileServer(http.Dir(toServe)))

	http.ListenAndServe(":1810", nil)
}

func ok(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func randomNum(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.FormatInt(rand.Int63(), 10)))
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
