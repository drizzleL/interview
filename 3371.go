package main

import "sort"

func getLargestOutlier(nums []int) int {
	var sum int
	dict := map[int]int{}
	for _, num := range nums {
		dict[num] += 1
		sum += num
	}
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		v := nums[i]
		tmpSum := sum - v
		if tmpSum%2 != 0 {
			continue
		}
		if dict[tmpSum/2] == 0 {
			continue
		}
		if dict[tmpSum/2] >= 2 || tmpSum/2 != v {
			return v
		}
	}
	return 0
}
