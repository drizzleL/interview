package main

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	visited := make([]bool, len(nums))
	var helper func(curr []int)
	helper = func(curr []int) {
		if len(curr) == len(nums) {
			ret = append(ret, append([]int{}, curr...))
			return
		}
		for i := range nums {
			if visited[i] {
				continue
			}
			if i != 0 && nums[i-1] == nums[i] && !visited[i-1] {
				continue
			}
			visited[i] = true
			helper(append(curr, nums[i]))
			visited[i] = false
		}
	}
	helper(nil)
	return ret
}
