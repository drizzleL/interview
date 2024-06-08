package main

func minLengthAfterRemovals(nums []int) int {
	dict := map[int]int{}
	var maxSum int
	for _, num := range nums {
		dict[num]++
		maxSum = max(maxSum, dict[num])
	}
	if maxSum <= len(nums)/2 {
		return len(nums) % 2
	}
	return maxSum*2 - len(nums)
}
