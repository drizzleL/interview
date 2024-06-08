package main

import (
	"sort"
)

func maxSumDivThree(nums []int) int {
	sort.Ints(nums)
	var onecnt, twocnt []int
	var ret int
	for _, num := range nums {
		switch num % 3 {
		case 0:
			ret += num
		case 1:
			onecnt = append(onecnt, num)
		case 2:
			twocnt = append(twocnt, num)
		}
	}
	helper := func(a, b []int) int {
		var ret int
		for len(a)%3 != 0 && len(b) != 0 {
			ret += a[len(a)-1] + b[len(b)-1]
			a = a[:len(a)-1]
			b = b[:len(b)-1]
		}
		for len(a) >= 3 {
			ret += a[len(a)-1] + a[len(a)-2] + a[len(a)-3]
			a = a[:len(a)-3]
		}
		for len(b) >= 3 {
			ret += b[len(b)-1] + b[len(b)-2] + b[len(b)-3]
			b = b[:len(b)-3]
		}
		return ret
	}
	ret += max(helper(onecnt, twocnt), helper(twocnt, onecnt))
	return ret
}
