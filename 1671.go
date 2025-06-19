package main

import (
	"math"
	"sort"
)

func minimumMountainRemovals(nums []int) int {
	lis := func(nums []int) []int {
		var q []int
		ret := make([]int, len(nums))
		for i, num := range nums {
			idx := sort.SearchInts(q, num)
			if idx == len(q) {
				q = append(q, num)
			}
			q[idx] = num
			ret[i] = idx + 1
		}
		return ret
	}
	rev := func(x []int) {
		for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
			x[i], x[j] = x[j], x[i]
		}
	}
	inc := lis(nums)
	rev(nums)
	dec := lis(nums)
	rev(dec)
	ret := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if inc[i] > 1 && dec[i] > 1 {
			ret = min(ret, len(nums)+1-inc[i]-dec[i])
		}
	}
	return ret
}
