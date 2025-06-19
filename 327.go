package main

import (
	"sort"
)

func countRangeSum(nums []int, lower int, upper int) int {
	sums := make([]int, len(nums)+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}
	var helper func(l, r int) int
	helper = func(l, r int) int {
		mid := (l + r) / 2
		if mid == l {
			return 0
		}
		ret := helper(l, mid) + helper(mid, r)
		i, j := mid, mid
		for _, v := range sums[l:mid] {
			for i < r && sums[i]-v < lower {
				i++
			}
			for j < r && sums[j]-v <= upper {
				j++
			}
			ret += j - i
		}
		sort.Ints(sums[l:r])
		return ret
	}
	return helper(0, len(sums))
}
