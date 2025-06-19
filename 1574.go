package main

import "sort"

func findLengthOfShortestSubarray(arr []int) int {
	var j int
	for ; j < len(arr)-1 && arr[j] <= arr[j+1]; j++ {
	}
	if j == len(arr)-1 {
		return 0
	}
	m := len(arr) - 1
	for ; arr[m] >= arr[m-1]; m-- {
	}
	ret := m
	for i := 0; i <= j; i++ {
		idx := sort.SearchInts(arr[m:], arr[i])
		ret = min(ret, m+idx-i-1)
	}
	return ret
}
