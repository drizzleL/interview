package main

func maxAdjacentDistance(nums []int) int {
	var ret int
	for i := 0; i < len(nums); i++ {
		next := (i + 1) % len(nums)
		ret = max(ret, abs(nums[i]-nums[next]))
	}
	return ret
}
