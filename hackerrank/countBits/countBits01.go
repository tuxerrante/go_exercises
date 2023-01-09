package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

/*
* implement a function that counts the number of set of bits in binary representation
* Complete the 'countBits' function below.
* The function is expected to return an int32.
* The function accepts uint32 num as parameter.

* https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
*
 */
func countBitsString(num uint32) int32 {

	num_64 := int64(num)
	num_2 := strconv.FormatInt(num_64, 2)
	_, counter := recursiveCount(num_2, 0)
	return counter
}

func recursiveCount(s string, counter int32) (string, int32) {
	if len(s) < 1 {
		return "", counter
	}

	if s[:1] == "1" {
		counter += 1
	}
	//fmt.Printf("Current counter: %d\n", counter)
	//fmt.Printf("Current char %s\n", s[:1])

	return recursiveCount(s[1:], counter)
}

func countBitsBinary(num uint32) int32 {
	// for binary_number > 0
	//		if current bit == 1 => counter++
	//		shift left binary_number
	var counter uint32 = 0
	for num > 0 {
		counter += num & 1
		num >>= 1
	}
	return int32(counter)
}

func countBitsSmart(num uint32) int32 {
	var upBits int32
	var highestPower uint32 = math.MaxUint32

	if num%2 > 0 {
		upBits += 1
		// fmt.Printf("\tUP bits = %d\n", upBits)
	}
	// upBits = 1 because is odd
	for ; num > 2 && highestPower > 1; upBits++ { // num  = 81.           | num = 17; upBits = 2  | num = 3 upBits = 3
		highestPower = uint32(math.Log2(float64(num)))         // log2(81) = 6 		| log2(17)  = 3 	    | log2(3) = 1
		num = num - uint32(math.Pow(2, float64(highestPower))) // num    = 81-64 = 17	| num    = 17-8 = 9     | num = 1
		// fmt.Printf("\tlog2 = %d\n", highestPower)
		// fmt.Printf("\tnum %d rest %d\n", num, num%2)
		// fmt.Printf("\tUP bits = %d\n", upBits)
	}

	return upBits
}

// Profiling with execution time and calling
func invoker(numInput []uint32, funcName string, countFunc func(num uint32) int32) {
	fmt.Println("\n=========================================================")
	start := time.Now()

	for _, numTemp := range numInput {
		// fmt.Printf("> Number:  %d\n", numTemp)
		// fmt.Printf("\t%s\n", strconv.FormatInt(int64(numTemp), 2))
		// fmt.Printf("> Number:  %d (%s)\tUp bits: %d\n", numTemp, strconv.FormatInt(int64(numTemp), 2), countFunc( uint32(numTemp) ))
		countFunc(uint32(numTemp))
	}
	fmt.Printf("Time elapsed for %v func= %vs", funcName, time.Since(start))
}

func main() {
	const testSize = 10000000
	var numInput = make([]uint32, testSize)
	for i := 0; i < testSize; i++ {
		numInput[i] = uint32(rand.Intn(500))
	}

	fmt.Printf("\n Test size = %d\n", testSize)

	// No need to use pointers when using slices instead of array https://stackoverflow.com/a/38732782/3673430
	invoker(numInput, "countBitsBinary", countBitsBinary)
	invoker(numInput, "countBitsString", countBitsString)
	invoker(numInput, "countBitsSmart", countBitsSmart)

	fmt.Println()
}
