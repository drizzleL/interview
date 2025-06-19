package main

import (
	"sort"
)

func maxFrequencyScore(nums []int, k int64) int {
	sort.Ints(nums)
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	getPre := func(l, r int) int {
		if l > r {
			return 0
		}
		return preSum[r+1] - preSum[l]
	}
	check := func(l, r int) bool {
		mid := (l + r) / 2
		val := nums[mid]
		size1, size2 := mid-l, r-mid
		cost := (size1*val - getPre(l, mid-1)) + (getPre(mid+1, r) - size2*val)
		return cost <= int(k)
	}
	var ret int
	for l, r := 0, 0; l < len(nums); l++ {
		for ; r < len(nums) && check(l, r); r++ {
		}
		ret = max(ret, r-l)
	}
	return ret
}
