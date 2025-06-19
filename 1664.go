package main

func waysToMakeFair(nums []int) int {
	gaps := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		gaps[i+1] = gaps[i]
		if i%2 == 1 {
			gaps[i+1] += nums[i]
		} else {
			gaps[i+1] -= nums[i]
		}
	}
	var ret int
	for i := 0; i < len(nums); i++ {
		beforeGap := gaps[i]
		afterGap := gaps[len(gaps)-1] - gaps[i+1]
		if beforeGap == afterGap {
			ret += 1
		}
	}
	return ret
}
