package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 0 {
		return 0
	}
	pre := 1
	last := 0
	var ret int
	for i, num := range nums {
		if num >= k {
			pre = 1
			last = i + 1
			continue
		}
		pre *= num
		for pre >= k {
			pre /= nums[last]
			last += 1
		}
		ret += i - last + 1
	}
	return ret
}
