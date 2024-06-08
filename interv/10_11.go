package main

import "sort"

func wiggleSort2(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for (i%2 == 1 && nums[i] < nums[i+1]) || (i%2 == 0 && nums[i] > nums[i+1]) {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
}

func wiggleSort(nums []int) {
	sort.Ints(nums)
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	l, r := 0, len(nums)-1
	for i := 0; i < len(nums); i += 2 {
		nums[i] = tmp[r]
		r--
		if i+1 < len(nums) {
			nums[i+1] = tmp[l]
			l++
		}
	}
}
