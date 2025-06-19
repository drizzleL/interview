package main

import "sort"

func maxStrength(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}
	var negs []int
	var posCnt int
	pos := 1
	for _, num := range nums {
		switch {
		case num > 0:
			pos *= num
			posCnt += 1
		case num < 0:
			negs = append(negs, num)
		}
	}
	neg := 1
	sort.Ints(negs)
	for i := 0; i+1 < len(negs); i += 2 {
		neg *= negs[i]
		neg *= negs[i+1]
	}
	if posCnt > 0 {
		return int64(pos * neg)
	}
	if len(negs) >= 2 { // all negs
		return int64(neg)
	}
	return 0
}
