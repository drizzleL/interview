package main

func maximumUniqueSubarray(nums []int) int {
	var ret int
	var sum int
	seen := map[int]bool{}
	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]
		for seen[nums[j]] {
			sum -= nums[i]
			delete(seen, nums[i])
			i++
		}
		seen[nums[j]] = true
		ret = max(ret, sum)
	}
	return ret
}
