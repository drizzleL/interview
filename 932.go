package main

func beautifulArray(n int) []int {
	dp := make([][]int, n+1)
	dp[1] = []int{1}
	getCpy := func(original []int) []int {
		ret := make([]int, len(original))
		copy(ret, original)
		return ret
	}
	for i := 2; i <= n; i++ {
		left, right := getCpy(dp[(i+1)/2]), getCpy(dp[i/2])
		for i := range left {
			left[i] = left[i]*2 - 1
		}
		for i := range right {
			right[i] = right[i] * 2
		}
		dp[i] = append(left, right...)
	}
	return dp[n]
}
