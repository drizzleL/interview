package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret, tmp := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		tmp = max(num, tmp+num)
		ret = max(ret, tmp)
	}
	return ret
}
