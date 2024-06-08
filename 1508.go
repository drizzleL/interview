package main

import "sort"

func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		var tmp int
		for j := i; j < n; j++ {
			tmp += nums[j]
			tmp %= 1e9 + 7
			sums = append(sums, tmp)
		}
	}
	sort.Ints(sums)
	var ret int
	for i := left - 1; i < right; i++ {
		ret += sums[i]
		ret %= 1e9 + 7
	}
	return ret
}
