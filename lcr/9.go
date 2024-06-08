package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	var ret int
	if k <= 1 {
		return 0
	}
	for i, j, curr := 0, 0, 1; i < len(nums); i++ {
		curr *= nums[i]
		for curr >= k {
			curr /= nums[j]
			j++
		}
		ret += i - j + 1
	}
	return ret
}
