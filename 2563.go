package main

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var l, r int
	var ret int
	for i := range nums {
		for l < len(nums) && nums[i]+nums[l] < lower {
			l += 1
		}
		for r < len(nums) && nums[i]+nums[r] <= upper {
			r += 1
		}
		ret += r - l
	}
	return int64(ret)
}
func countFairPairs2(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var ret int
	for i, num := range nums {
		a := sort.SearchInts(nums, lower-num)
		if a >= len(nums) {
			continue
		}
		a = max(a, i+1)
		b := sort.SearchInts(nums, upper-num+1)
		if b <= i {
			continue
		}
		ret += b - a
	}
	return int64(ret)
}
