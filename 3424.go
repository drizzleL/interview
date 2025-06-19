package main

import "sort"

func minCost4(arr []int, brr []int, k int64) int64 {
	helper := func() int {
		var ret int
		for i := range arr {
			ret += abs(arr[i] - brr[i])
		}
		return ret
	}
	notChange := helper()
	sort.Ints(arr)
	sort.Ints(brr)
	changed := helper() + int(k)
	return int64(min(changed, notChange))
}
