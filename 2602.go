package main

import "sort"

func minOperations19(nums []int, queries []int) []int64 {
	sort.Ints(nums)
	presum := make([]int, len(nums)+1)
	for i := 1; i < len(presum); i++ {
		presum[i] = presum[i-1] + nums[i-1]
	}
	getSum := func(i, j int) int {
		if j < 0 {
			return 0
		}
		if i == len(nums) {
			return 0
		}
		return presum[j+1] - presum[i]
	}
	ret := make([]int64, len(queries))
	for i, q := range queries {
		idx := sort.SearchInts(nums, q)
		size1, size2 := idx, len(nums)-idx
		ret[i] = int64((size1*q - getSum(0, idx-1)) + (getSum(idx, len(nums)-1) - size2*q))
	}
	return ret
}
