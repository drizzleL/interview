package main

func bowlSubarrays(nums []int) int64 {
	left, right := make([]bool, len(nums)), make([]bool, len(nums))
	leftMax, rightMax := nums[0], nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		if leftMax > nums[i] {
			left[i] = true
		} else {
			leftMax = nums[i]
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if rightMax > nums[i] {
			right[i] = true
		} else {
			rightMax = nums[i]
		}
	}
	var ret int
	for i := 0; i < len(nums); i++ {
		if left[i] && right[i] {
			ret += 1
		}
	}
	return int64(ret)
}
