package main

import "sort"

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	var sum int
	for i, num := range nums {
		sum += num
		nums[i] = num ^ k - nums[i]
	}
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 1; i -= 2 {
		tmp := nums[i] + nums[i-1]
		if tmp <= 0 {
			break
		}
		sum += tmp
	}
	return int64(sum)
}
