package basicExamples

import (
	"errors"
	"fmt"
	"net/http"
)

func BasicExamples() error {

	var ok error

	ok = httpExample()

	return ok
}

/*
Propagate an error
https://pkg.go.dev/net/http#Client.Get
*/
func httpExample() error {
	const url string = "http://httpstat.us/200"

	// An error is returned if there were too many redirects or if there was an HTTP protocol error. A non-2xx response doesn't cause an error.
	http_resp, http_err := http.Get(url)

	http_err = errors.New("Fake HTTP error!")
	if http_err != nil {
		// If the format specifier includes a %w verb with an error operand, the returned error will implement an Unwrap method
		return fmt.Errorf("HTTP get url: %w", http_err)
	}

	// body, err := io.ReadAll(http_resp.Body)
	// defer body.Body.Close()

	fmt.Println(http_resp.StatusCode)

	return nil
}

/*
func logExample(){

	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}
*/
