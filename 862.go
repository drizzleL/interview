package main

import "math"

func shortestSubarray(nums []int, k int) int {
	ret := math.MaxInt32
	var q []int
	for i := range nums {
		if i != 0 {
			nums[i] += nums[i+1]
		}
		if nums[i] >= k {
			ret = min(ret, i+1)
		}
		for len(q) != 0 && nums[i]-nums[q[0]] >= k {
			ret = min(ret, i-q[0])
			q = q[1:]
		}
		for len(q) != 0 && nums[i] <= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
	}
	if ret == math.MaxInt32 {
		return -1
	}
	return ret
}
