package leetcode

import (
	"fmt"
	"unsafe"
)

func runningSum(nums []int) []int {

	result := make([]int, len(nums))

	// 10^6 >? size(int)
	myint_32 := int32(1)
	myint := int(1)
	myint_64 := int64(1)

	// reflect Type.Size() returns the number of bytes needed to store
	// a value of the given type; it is analogous to unsafe.Sizeof.
	fmt.Printf("%T = %d bytes\n", myint, unsafe.Sizeof(myint))
	fmt.Printf("%T = %d bytes\n", myint_32, unsafe.Sizeof(myint_32))
	fmt.Printf("%T = %d bytes\n", myint_64, unsafe.Sizeof(myint_64))

	// int here is int64 = 8 bytes = 9*10^18
	// max value could be an array of 1000 times 10^6

	if len(nums) <= 0 {
		return nil
	}

	if len(nums) >= 1 {
		result[0] = nums[0]
	}

	for i, num_value := range nums[1:] {
		i += 1
		result[i] = result[i-1] + num_value
		fmt.Printf("result[i] %d = result[i-1] %d + num_value %d\n", result[i], result[i-1], num_value)
	}

	return result
}
