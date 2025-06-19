package main

import (
	"fmt"
	"sort"
)

func minCost6(nums []int) int {
	cache := map[string]int{}
	var helper func(i int, preVal int) int
	helper = func(i int, preVal int) (ret int) {
		key := fmt.Sprintf("%d_%d", i, preVal)
		if c, ok := cache[key]; ok {
			return c
		}
		defer func() {
			cache[key] = ret
		}()
		if i == len(nums) { // last one
			return preVal
		}
		if i == len(nums)-1 { // last two
			return max(preVal, nums[i])
		}
		a, b := nums[i], nums[i+1]
		vals := []int{preVal, a, b}
		sort.Ints(vals)
		ret = helper(i+2, vals[0]) + vals[2]
		if vals[1] != vals[2] {
			ret = min(ret, helper(i+2, vals[2])+vals[1])
		}
		return
	}
	return helper(1, nums[0])
}
