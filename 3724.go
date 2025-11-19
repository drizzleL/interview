package main

import "math"

func minOperations37(nums1 []int, nums2 []int) int64 {
	var ret int
	extra := math.MaxInt32
	val := nums2[len(nums1)]
	for i := 0; i < len(nums1); i++ {
		a, b := min(nums1[i], nums2[i]), max(nums1[i], nums2[i])
		ret += b - a
		if val >= a && val <= b {
			extra = 0
		} else if val < a {
			extra = min(extra, a-val)
		} else {
			extra = min(extra, val-b)
		}
	}
	return int64(ret + extra + 1)
}
