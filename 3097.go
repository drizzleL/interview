package main

import (
	"math"
)

func minimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}
	var i int
	dict := [32]int{}
	var v int
	ret := math.MaxInt32
	for j := 0; j < len(nums); j++ {
		for m := 0; m < 32; m++ {
			dict[m] += (nums[j] & (1 << m)) * 1
		}
		v |= nums[j]
		for v >= k {
			ret = min(ret, j-i+1)
			for m := 0; m < 32; m++ {
				dict[m] -= (nums[i] & (1 << m)) * 1
				if dict[m] == 0 {
					v &= ^(1 << m)
				}
			}
			i++
		}
	}
	if ret == math.MaxInt32 {
		return -1
	}
	return ret
}
