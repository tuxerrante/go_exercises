package main

import (
	"fmt"
	"math"
)

func main() {

	const overflower uint8 = 100
	fmt.Println("Warning number = ", math.Pow(2, 8)-1)

	var a, b uint8 = 0, 0

}

func imdangerous(x int8) {

	for ; a < overflower; a, b = a+1, b+1 {
		var c = float64(a) * b
		fmt.Println(a, b, c)
	}
}
