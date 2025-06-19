package main

import "sort"

func sortEvenOdd(nums []int) []int {
	var odds, evens []int
	for i := 0; i < len(nums); i += 2 {
		evens = append(evens, nums[i])
		if i+1 < len(nums) {
			odds = append(odds, nums[i+1])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(odds)))
	sort.Ints(evens)
	var ret []int
	for i := 0; i < len(evens); i += 1 {
		ret = append(ret, evens[i])
		if i < len(odds) {
			ret = append(ret, odds[i])
		}
	}
	return ret
}
