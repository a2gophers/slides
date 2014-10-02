package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		fmt.Println("hit ok")
	})
	http.ListenAndServe(":1809", nil)
}
