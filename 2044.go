package main

import "math"

func countMaxOrSubsets(nums []int) int {
	var or int
	for _, num := range nums {
		or |= num
	}
	var ret int
	var helper func(i int, flag int)
	helper = func(i int, flag int) {
		if i == len(nums) {
			if flag == or {
				ret += 1
			}
			return
		}
		if flag == or {
			ret += int(math.Pow(2, float64(len(nums)-i)))
			return
		}
		helper(i+1, flag|nums[i])
		helper(i+1, flag)
	}
	helper(0, 0)
	return ret
}
