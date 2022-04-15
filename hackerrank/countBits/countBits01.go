package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"math"
)

/*
 * implement a function that counts the number of set of bits in binary representation
 * Complete the 'countBits' function below.
 * The function is expected to return an int32.
 * The function accepts unit32 num as parameter.
*/
func countBitsString(num uint32) int32 {

	num_64 := int64(num)

	num_2 := strconv.FormatInt(num_64, 2)
	fmt.Printf("\t%s\n", num_2)

	_, counter := recursiveCount(num_2, 0)
	return counter
}

func recursiveCount(s string, counter int32) (string, int32) {
	if len(s) < 1 {
		return "", counter
		}
	
	if s[:1] == "1" {
		counter+=1
	}
	//fmt.Printf("Current counter: %d\n", counter)
	//fmt.Printf("Current char %s\n", s[:1])	
	
	return recursiveCount(s[1:], counter)
}


func countBitsBinary(num uint32) int32 {
	fmt.Println("TODO")
	// convert uint23 in binary form
	// for binary_number > 0 
	//		if current bit == 1 => counter++
	//		shift left binary_number
	return 0
}

func countBitsSmart(num uint32) uint32{
	var upBits 			uint32
	var highestPower	uint32 = math.MaxUint32

	if num % 2 >0  {
		upBits+=1
		fmt.Printf("\tUP bits = %d\n", upBits)
	}
																			// upBits = 1 because is odd
	for ; (num > 2 && highestPower > 1); upBits++ {							// num  = 81.           | num = 17; upBits = 2  | num = 3 upBits = 3
		// highest power of 2
		highestPower = uint32(math.Log2(float64(num)))						// log2(81) = 6 		| log2(17)  = 3 	| log2(3) = 1
		num          = num - uint32(math.Pow(2, float64(highestPower))) 	// num    = 81-64 = 17	| num    = 17-8 = 9 | num = 1
		fmt.Printf("\tlog2 = %d\n", highestPower)							
		fmt.Printf("\tnum %d rest %d\n", num, num%2)
		fmt.Printf("\tUP bits = %d\n", upBits)
	}
	
	return upBits
}


func main() {
	const testSize = 1000000
	var numInput   = [testSize]uint32{}
	for i:=0; i<testSize; i++ {
		numInput[i] = uint32(rand.Intn(500))
	}
	
	fmt.Println("=========================================================")
	for _, numTemp := range numInput{
		fmt.Printf("> Number:  %d\n", numTemp)
		num := uint32(numTemp)
		fmt.Printf("\tUp bits: %d\n", countBitsString(num))
	}
	
	fmt.Println("=========================================================")
	for _, numTemp := range numInput{
		fmt.Printf("> Number:  %d\n", numTemp)
		num := uint32(numTemp)
		fmt.Printf("\tUp bits: %d\n", countBitsSmart(num))
	}
}
