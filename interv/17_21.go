package main

func trap(height []int) int {
	lh, rh := make([]int, len(height)), make([]int, len(height))
	var tmp int
	for i := 0; i < len(height); i++ {
		tmp = max(tmp, height[i])
		lh[i] = tmp
	}
	tmp = 0
	for i := len(height) - 1; i >= 0; i-- {
		tmp = max(tmp, height[i])
		rh[i] = tmp
	}
	var ret int
	for i := 0; i < len(height); i++ {
		ret += max(lh[i], rh[i]) - height[i]
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
