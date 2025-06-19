package main

import "sort"

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	var ret int
	var reduce int
	nums1Cpy := make([]int, len(nums1))
	copy(nums1Cpy, nums1)
	sort.Ints(nums1Cpy)
	getMin := func(x int) int {
		idx := sort.SearchInts(nums1Cpy, x)
		if idx == 0 {
			return nums1Cpy[idx] - x
		}
		if idx == len(nums1Cpy) {
			return x - nums1Cpy[idx-1]
		}
		return min(nums1Cpy[idx]-x, x-nums1Cpy[idx-1])
	}
	for i := range nums1 {
		diff := abs(nums1[i] - nums2[i])
		ret += diff
		ret %= 1000000007
		reduce = max(reduce, abs(diff-getMin(nums2[i])))
	}
	ret -= reduce
	if ret < 0 {
		ret += 1e9 + 7
	}
	return ret
}
