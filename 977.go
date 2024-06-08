package main

func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	ret := make([]int, len(nums))
	i := len(nums) - 1
	for l <= r {
		if nums[l]*nums[l] <= nums[r]*nums[r] {
			ret[i] = nums[r] * nums[r]
			r--
		} else {
			ret[i] = nums[l] * nums[l]
			l++
		}
		i--
	}
	return ret
}
