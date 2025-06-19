package main

func splitArraySameAverage(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum == 0 {
		return true
	}
	dp := make([]int, sum)
	dp[0] = 1
	check := func(sum1, size1 int) bool {
		sum2 := sum - sum1
		size2 := len(nums) - size1
		return sum1*size2 == sum2*size1
	}
	for _, num := range nums {
		for sum1 := sum - 1; sum1 >= num; sum1-- {
			dp[sum1] |= dp[sum1-num] << 1
		}
	}
	for sum1, mask := range dp {
		for size1 := 1; size1 < len(nums); size1++ {
			if 1<<size1&mask == 0 {
				continue
			}
			if check(sum1, size1) {
				return true
			}
		}
	}
	return false
}
