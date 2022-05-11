package main

import (
	"fmt"
	"myErrors/basicExamples"
)

func main() {

	fmt.Println("===== ERROR PACKAGE =====")
	/* 	
		https://pkg.go.dev/errors#pkg-overview
		https://pkg.go.dev/builtin#error

		The built-in type error is an interface type. 
		a non-nil error has an error message string which we can obtain by calling its Error method or print by calling
		fmt.Println(err) or fmt.Printf("%v", err).
	
		Go programs use ordinary control-flow mechanisms like if and return to respond to errors. 
		This style undeniably demands that more attention be paid to error-handling logic, but that is precisely the point!
	*/

	err := basicExamples.BasicExamples()
	if err != nil {
		// panic(err)
		// os.Exit(1)
		fmt.Printf("%v\n", err.Error())
	}
}