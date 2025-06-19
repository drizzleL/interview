package main

import (
	"math"
)

func minMove3(nums []int, limit int) int {
	if limit == 1 {
		return 0
	}
	diff := make([]int, limit*2+2)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		a, b := nums[i], nums[j]
		if a > b {
			a, b = b, a
		}
		diff[2] += 2
		diff[min(a, b)+1] -= 1
		diff[a+b] -= 1
		diff[a+b+1] += 1
		diff[max(a, b)+limit+1] += 1
	}
	ret := math.MaxInt32
	var v int
	for i := 2; i <= limit*2; i++ {
		v += diff[i]
		ret = min(ret, v)
	}
	return ret
}
