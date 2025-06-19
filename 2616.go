package main

import "sort"

func minimizeMax(nums []int, p int) int {
	var l, r int
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		r = max(r, nums[i]-nums[i-1])
	}
	check := func(x int) bool {
		var cnt int
		for i := 1; i < len(nums); {
			if nums[i]-nums[i-1] <= x {
				i += 2
				cnt += 1
			} else {
				i += 1
			}
		}
		return cnt >= p
	}
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
