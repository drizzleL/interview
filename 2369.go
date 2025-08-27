package main

func validPartition(nums []int) bool {
	dp := make([]bool, len(nums)+1)
	dp[len(nums)] = true
	check1 := func(i int) bool {
		if i+1 >= len(nums) {
			return false
		}
		return nums[i] == nums[i+1] && dp[i+2]
	}
	check2 := func(i int) bool {
		if i+2 >= len(nums) {
			return false
		}
		return nums[i] == nums[i+1] && nums[i] == nums[i+2] && dp[i+3]
	}
	check3 := func(i int) bool {
		if i+2 >= len(nums) {
			return false
		}
		return nums[i]+1 == nums[i+1] && nums[i+1]+1 == nums[i+2] && dp[i+3]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = check1(i) || check2(i) || check3(i)
	}
	return dp[0]
}
