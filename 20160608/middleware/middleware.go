package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

func main() {

	// START OMIT
	c := xhandler.Chain{}
	c.UseC(func(next xhandler.HandlerC) xhandler.HandlerC {
		return xhandler.HandlerFuncC(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL.Path)
			next.ServeHTTPC(ctx, w, r)
		})

	})

	mux := xmux.New()
	mux.GET("/yolo/:bolo", xhandler.HandlerFuncC(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(xmux.Param(ctx, "name")))
	}))

	fmt.Println("listening on :18016...")
	fmt.Println(http.ListenAndServe(":18016", c.Handler(mux)))
	// END OMIT
}
