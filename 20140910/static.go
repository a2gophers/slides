package main

import (
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

	http.Handle("/", http.FileServer(http.Dir(toServe)))

	http.ListenAndServe(":1810", nil)
}
