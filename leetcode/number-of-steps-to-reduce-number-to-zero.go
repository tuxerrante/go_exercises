package leetcode

/*
   https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
   Given an integer num, return the number of steps to reduce it to zero.
   In one step, if the current number is even, you have to divide it by 2,
   otherwise, you have to subtract 1 from it.

   INPUT:
   0 <= num <= 10^6
*/
import (
	"encoding/binary"
	// "fmt"
	"math"
	// "testing"
)

func numberOfSteps(num int) int {
	const maxNumOfBytes int = 4

	// return myWay(num, maxNumOfBytes)
	return smartWay(uint32(num))
}

/*
	    Convert number from decimal to [4]byte
		Iterate over the least important bit
		  if it is 1: shift left and increase sum by 2
		  if it is 0: shift left and incresse sum by 1
*/
func myWay(num, maxNumOfBytes int) int {
	steps, index := 0, 0
	numBytes := make([]byte, maxNumOfBytes)

	binary.LittleEndian.PutUint32(numBytes, uint32(num))
	currentByte := numBytes[index]
	// fmt.Printf("Input %d\n", num)
	// fmt.Printf("%b\n", numBytes)

	if num == 0 {
		return 0
	}

	// When I find a 1 on a byte I should count the steps done for the previous byte
	for index < maxNumOfBytes {
		// fmt.Printf(">> %b.\n", currentByte)
		// Go next until there is a 1
		if currentByte == 0 {
			index++
			if index < 4 {
				currentByte = numBytes[index]
			}
			continue
		} else if currentByte%2 != 0 {
			currentByte -= 0x1
			steps += 1
		}
		// Divide by 2
		currentByte >>= 1
		// fmt.Printf(">>  Divided by 2.\n    Next byte=%b, steps=%d\n", currentByte, steps)
	}

	steps += int(math.Log2(float64(num)))
	// fmt.Printf(">> Summing all previous bits: %d\n", int(math.Log2(float64(num))))

	return steps
}

/*
https://stackoverflow.com/a/109025/3673430
https://cs.opensource.google/go/go/+/refs/tags/go1.21.3:src/math/bits/bits.go;l=135
*/
func smartWay(num uint32) int {
	if num == 0 {
		return 0
	}
	return countOnes(num)
}

func countOnes(byteNumber uint32) int {
	count := 1
	for byteNumber > 1 {
		// fmt.Printf("> %b\n", byteNumber)
		// Sum every time there is a 1 + the division by 2
		count += int(byteNumber&1) + 1
		byteNumber >>= 1
	}
	return count
}

// func TestNumberOfSteps(t *testing.T) {
// 	in := []int{0, 1, 2, 3, 14, 123, 256, 513, 83962}
// 	want := []int{0, 1, 2, 3, 6, 12, 9, 11, 27}

// 	for tc, _ := range in {
// 		t.Run(fmt.Sprintf("%d", in[tc]), func(t *testing.T) {
// 			// t.Parallel()
// 			got := numberOfSteps(in[tc])
// 			if want[tc] != got {
// 				t.Errorf("Got = %d. Want = %d", got, want[tc])
// 			}
// 		})
// 	}
// }
