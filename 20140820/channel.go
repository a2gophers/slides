// +build OMIT
package main

import "fmt"

func main() {
	// START OMIT
	// make a channel
	c1 := make(chan bool)

	// send something over a channel
	c <- true

	// receive something from a channel
	a := <-c
	fmt.Println("a is: ", a)

	// channels can send anything
	funcChan := make(chan func(string))
	funcChan <- func(s string) { fmt.Println(s) }
	f := <-funcChan
	f("Hello, gophers!")

	// STOP OMIT
}
