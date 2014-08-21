package main

import "fmt"

type info struct {
	name string
}

func main() {
	infoChan, doneChan := make(chan info), make(chan string)

	go func() {
		<-infoChan // this will block
		doneChan <- "info closure"
	}()

	infoChan <- info{name: "unbuffered channel!"}

	fmt.Printf("%q just finished\n", <-doneChan)
}
