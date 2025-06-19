package main

func maxArrayValue(nums []int) int64 {
	curr := nums[len(nums)-1]
	ret := curr
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= curr {
			curr += nums[i]
		} else {
			curr = nums[i]
		}
		ret = max(ret, curr)
	}
	return int64(ret)
}
