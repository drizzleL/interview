package main

import "sort"

func maxWidthRamp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lastIdx := map[int]int{}
	for i, num := range nums {
		lastIdx[num] = i
	}
	s := make([]int, 0, len(lastIdx))
	for k := range lastIdx {
		s = append(s, k)
	}
	sort.Ints(s)
	dict := map[int]int{}
	var ret int
	for i := len(s) - 1; i >= 0; i-- {
		dict[s[i]] = lastIdx[s[i]]
		if i != len(s)-1 {
			dict[s[i]] = max(dict[s[i]], dict[s[i+1]])
		}
	}
	for i, num := range nums {
		ret = max(ret, dict[num]-i)
	}
	return ret
}
