package main

import "sort"

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return nil
		}
		ret = append(ret, []int{nums[i], nums[i+1], nums[i+2]})
	}
	return ret
}
