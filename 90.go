package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	visited := make([]bool, len(nums))
	var ret [][]int
	var helper func(i int, curr []int)
	helper = func(i int, curr []int) {
		if i == len(nums) {
			ret = append(ret, append([]int{}, curr...))
			return
		}
		helper(i+1, curr)
		if i == 0 || nums[i-1] != nums[i] || visited[i-1] {
			visited[i] = true
			helper(i+1, append(curr, nums[i]))
			visited[i] = false
		}
	}
	helper(0, nil)
	return ret
}
