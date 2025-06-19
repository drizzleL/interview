package main

func buildArray(nums []int) []int {
	ret := make([]int, len(nums))
	for i, num := range nums {
		ret[i] = nums[num]
	}
	return ret
}
