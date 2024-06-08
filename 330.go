package main

func minPatches(nums []int, n int) int {
	var x int
	var ret int
	for j := 0; x < n && j < len(nums); {
		if nums[j] <= x+1 {
			x += nums[j]
			j += 1
			continue
		}
		x = x*2 + 1
		ret += 1
	}
	for x < n {
		x = x*2 + 1
		ret += 1
	}
	return ret
}
