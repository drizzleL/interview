package main

import "sort"

func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
	dict := map[int]int{}
	ret := -1
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		lastCnt := dict[num*num]
		dict[num] = lastCnt + 1
		if lastCnt != 0 {
			ret = max(ret, dict[num])
		}
	}
	return ret
}
