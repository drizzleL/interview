package main

func getFinalState(nums []int, k int, multiplier int) []int {
	for i := 0; i < k; i++ {
		idx := -1
		for j, num := range nums {
			if idx == -1 || num < nums[idx] {
				idx = j
			}
		}
		nums[idx] *= multiplier
	}
	return nums
}
