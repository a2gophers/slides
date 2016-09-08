package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

func main() {

	mux := xmux.New()

	level0 := mux.NewGroup("/api")
	level1 := level0.NewGroup("/nested")
	level1.GET("/yolo/:bolo", xhandler.HandlerFuncC(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(xmux.Param(ctx, "name")))
	}))

	fmt.Println("listening on :18016...")
	fmt.Println(http.ListenAndServe(":18016", xhandler.New(context.Background(), mux)))
}
