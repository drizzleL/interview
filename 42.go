package main

func trap(height []int) int {
	lMaxIdx, rMaxIdx := -1, -1
	var totalSum int
	for i := 0; i < len(height); i++ {
		totalSum += height[i]
		if lMaxIdx == -1 || height[i] > height[lMaxIdx] {
			lMaxIdx = i
		}
	}
	for i := len(height) - 1; i >= 0; i-- {
		if rMaxIdx == -1 || height[i] > height[rMaxIdx] {
			rMaxIdx = i
		}
	}
	var sum int
	var tmp int
	for i := 0; i < lMaxIdx; i++ {
		tmp = max(tmp, height[i])
		sum += tmp
	}
	tmp = 0
	for i := len(height) - 1; i > rMaxIdx; i-- {
		tmp = max(tmp, height[i])
		sum += tmp
	}
	sum += height[lMaxIdx] * (rMaxIdx - lMaxIdx + 1)
	return sum - totalSum
}
