package main

import "sort"

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	var ret int
	for i, j := 0, len(nums)-1; i <= j; i++ {
		for ; j >= i && nums[i]+nums[j] > target; j-- {
		}
		if j < i {
			break
		}
		ret += fastPow(2, j-i, 1)
		ret %= 1e9 + 7
	}
	return ret
}
