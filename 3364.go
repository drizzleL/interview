package main

import "math"

func minimumSumSubarray(nums []int, l int, r int) int {
	sums := make([]int, len(nums)+1)
	for i := range nums {
		sums[i+1] = sums[i] + nums[i]
	}
	ret := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		for size := l; size <= r && i+size <= len(nums); size++ {
			sum := sums[i+size] - sums[i]
			if sum <= 0 {
				continue
			}
			ret = min(ret, sum)
		}
	}
	if ret == math.MaxInt32 {
		return -1
	}
	return ret
}
