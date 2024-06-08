package main

func minSubArrayLen(target int, nums []int) int {
	var ret int
	for i, j, sum := 0, 1, nums[0]; i < len(nums) && j <= len(nums); {
		if sum < target {
			if j < len(nums) {
				sum += nums[j]
			}
			j++
		} else {
			if ret == 0 || j-i < ret {
				ret = j - i
			}
			sum -= nums[i]
			i++
			if i == j {
				j++
			}
		}
	}
	return ret
}
