package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {

	mux := httprouter.New()
	mux.GET("/yolo/:bolo", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Write([]byte(ps.ByName("bolo")))
	})

	n := negroni.New()
	n.UseHandler(mux)

	fmt.Println("listening on :18016...")
	fmt.Println(http.ListenAndServe(":18016", n))
}
