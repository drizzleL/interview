package main

import "sort"

func putMarbles(weights []int, k int) int64 {
	var vals []int
	for i := 0; i < len(weights)-1; i++ {
		vals = append(vals, weights[i]+weights[i+1])
	}
	sort.Ints(vals)
	var ret int
	for i := 0; i < k; i++ {
		ret += vals[len(vals)-1-i] - vals[i]
	}
	return int64(ret)
}
