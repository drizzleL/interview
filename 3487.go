package main

func maxSum5(nums []int) int {
	seen := map[int]bool{}
	var sum int
	var foundPos bool
	maxNeg := nums[0]
	for _, num := range nums {
		if num < 0 {
			maxNeg = max(maxNeg, num)
		}
		if seen[num] {
			continue
		}
		foundPos = true
		sum += num
		seen[num] = true
	}
	if foundPos {
		return sum
	}
	return maxNeg
}
