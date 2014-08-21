// +build OMIT
package main

import (
	"fmt"
	"net/http"
)

func main() {
	var (
		work  = make(chan string, 500)
		valid = make(chan string, 500)
		urls  = []string{"http://github.com", "http://google.com"} // some long source
		done  = make(chan bool, 10)
	)

	var sendIfValid = func(
		url string,
		resp *http.Response,
		err error,
		valid chan string) {

		if err == nil && resp.StatusCode > 199 && resp.StatusCode < 300 {
			valid <- url
		}
	}

	// START OMIT
	worker := func(work chan string) {
		for url := range work {
			resp, err := http.Get(url)
			sendIfValid(url, resp, err, valid)
		}
		done <- true
	}

	for i := 0; i < 10; i++ {
		go worker(work)
	}

	for _, url := range urls {
		work <- url
	}
	close(work)

	for i := 0; i < 10; i++ {
		<-done
	}
	// STOP OMIT

	// START2 OMIT
	moreResults := true
	for moreResults {
		select {
		case v := <-valid:
			fmt.Println("valid url: ", v)
		default:
			fmt.Println("no more urls...exiting...")
			moreResults = false
		}
	}
	// STOP2 OMIT
}
