package leetcode

func twoSum(nums []int, target int) []int {
	// if nums[i] > target I can skip it
	// if the array is sorted and nums[i]>target I can return
	// but since a solution is guaranteed it will end before anyway:
	// otherwise in go1.21:
	//  slices.SortFunc(nums, func(a,b int){return a>b})

	// BRUTEFORCE
	// I can avoid checking on breaking since a solution is guaranteed
	// for k,v := range nums {
	//     for kk := k+1; kk < len(nums); kk++ {
	//         if (v + nums[kk] == target){
	//             return []int{k, kk}
	//         }
	//     }
	// }

	// HASHMAP
	numsMap := make(map[int]int, len(nums))
	for index, value := range nums {
		numsMap[value] = index
		difference := target - value

		// It should be checked wheter the numsMap index is the same as the current one.
		// If 'target-value' exists in the map, we've found the index of the second operand.
		secondIndex, ok := numsMap[difference]
		if ok {
			return []int{index, secondIndex}
		}
	}

	return nil
}
