package main

func subsetXORSum(nums []int) int {
	var ret int
	var helper func(i int, sum int)
	helper = func(i int, sum int) {
		if i == len(nums) {
			ret += sum
			return
		}
		helper(i+1, sum)
		helper(i+1, sum^nums[i])
	}
	helper(0, 0)
	return ret
}
