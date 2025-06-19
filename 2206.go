package main

import "sort"

func divideArray2(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i*2 < len(nums); i++ {
		if nums[i*2] != nums[i*2+1] {
			return false
		}
	}
	return true
}
