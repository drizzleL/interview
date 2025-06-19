package main

func maximumPossibleSize(nums []int) int {
	ret := len(nums)
	val := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= val {
			val = nums[i]
			continue
		}
		ret -= 1
	}
	return ret
}
