package main

func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	ret[0] = 1
	pre := 1
	for i := 0; i < len(nums); i++ {
		ret[i] = pre
		pre *= nums[i]
	}
	pre = 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] = ret[i] * pre
		pre *= nums[i]
	}
	return ret
}
