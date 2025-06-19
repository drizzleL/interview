package main

import "sort"

func increasingTriplet(nums []int) bool {
	var q []int
	for _, num := range nums {
		idx := sort.SearchInts(q, num+1)
		if idx == len(q) {
			q = append(q, 0)
		}
		q[idx] = num
		if len(q) == 3 {
			return true
		}
	}
	return false
}
