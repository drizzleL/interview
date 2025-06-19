package main

func countSubarrays3(nums []int) int {
	var ret int
	for i := 1; i < len(nums)-1; i++ {
		if (nums[i-1]+nums[i+1])*2 == nums[i] {
			ret += 1
		}
	}
	return ret
}
