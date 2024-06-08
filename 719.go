package main

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	sum := func(x int) int {
		var ret int
		for i, num := range nums {
			idx := sort.SearchInts(nums, num+x+1)
			ret += idx - i - 1
		}
		return ret
	}
	l, r := 0, nums[len(nums)-1]-nums[0]
	for l < r {
		mid := (l + r) / 2
		if sum(mid) < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
