package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var helper = func(nums []int, target int) (ret [][]int) {
		i, j := 0, len(nums)-1
		for i < j {
			sum := nums[i] + nums[j]
			if sum == target {
				ret = append(ret, []int{nums[i], nums[j]})
				for i+1 < j && nums[i+1] == nums[i] {
					i++
				}
				i++
				j--
			} else if sum > target {
				j--
			} else {
				i++
			}
		}
		return
	}
	var ret [][]int
	var i int
	for i < len(nums) {
		num := nums[i]
		for _, tp := range helper(nums[i+1:], -num) {
			ret = append(ret, []int{num, tp[0], tp[1]})
		}
		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i += 1
		}
		i += 1
	}
	return ret
}
