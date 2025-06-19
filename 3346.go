package main

import "sort"

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	var ret int
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		num := nums[i]
		sameSize := sort.SearchInts(nums, num+1) - i
		smaller := sort.SearchInts(nums, num-k)
		larger := sort.SearchInts(nums, num+k+1) - 1
		ret = max(ret, sameSize+min(larger-smaller+1-sameSize, numOperations))

		j := sort.SearchInts(nums, num+k*2+1) - 1
		ret = max(ret, min(j-i+1, numOperations))
	}
	return ret
}
