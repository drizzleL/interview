package main

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
	var loss int
	var ret int
	for i := len(happiness) - 1; k > 0; i, k = i-1, k-1 {
		if happiness[i]-loss <= 0 {
			break
		}
		ret += happiness[i] - loss
		loss += 1
	}
	return int64(ret)
}
