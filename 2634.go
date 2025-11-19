package main

func minOperations78(nums []int) int {
	var g, c int
	for _, num := range nums {
		if num == 1 {
			c += 1
		}
		g = gcd(g, num)
	}
	if c > 0 {
		return len(nums) - c
	}
	if g != 1 {
		return -1
	}
	minStep := len(nums)
	for i := 0; i < len(nums); i++ {
		var g int
		for j := i; j < len(nums); j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				minStep = min(minStep, j-i)
				break
			}
		}
	}
	return len(nums) + minStep - 1
}
