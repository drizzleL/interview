package main

import "sort"

func kIncreasing(arr []int, k int) int {
	lis := func(nums []int) int {
		var q []int
		for _, num := range nums {
			idx := sort.SearchInts(q, num+1)
			if idx == len(q) {
				q = append(q, num)
			} else {
				q[idx] = num
			}
		}
		return len(q)
	}
	var ret int
	for i := 0; i < k; i++ {
		var nums []int
		for j := i; j < len(arr); j += k {
			nums = append(nums, arr[j])
		}
		ret += len(nums) - lis(nums)
	}
	return ret
}
