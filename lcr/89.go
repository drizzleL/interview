package main

func rob(nums []int) int {
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		a, b = max(a, b), max(a+nums[i], b)
	}
	return max(a, b)
}
