package main

import "sort"

func lengthOfLIS(nums []int) int {
	var q []int
	for _, num := range nums {
		idx := sort.SearchInts(q, num)
		if idx == len(nums) {
			q = append(q, num)
			continue
		}
		q[idx] = num
	}
	return len(q)
}
