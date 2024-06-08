package main

import (
	"sort"
)

func sumDistance(nums []int, s string, d int) int {
	for i := range nums {
		var v int
		switch s[i] {
		case 'L':
			v = -1
		case 'R':
			v = 1
		}
		v *= d
		nums[i] += v
	}
	sort.Ints(nums)
	var ret int
	var curr int
	for i := 1; i < len(nums); i++ {
		gap := nums[i] - nums[i-1]
		curr = gap*i + curr
		ret += curr
	}
	return ret
}
