package main

func countHillValley(nums []int) int {
	var ret int
	var flag int
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		var newFlag int
		if nums[i] > nums[i-1] {
			newFlag = 1
		} else {
			newFlag = 2
		}
		if flag == 0 {
			flag = newFlag
			continue
		}
		if flag == newFlag {
			continue
		}
		ret += 1
		flag = newFlag
	}
	return ret
}
