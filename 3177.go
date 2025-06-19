package main

func maximumLength3(nums []int, k int) int {
	dp := make([]map[int]int, k+1)
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	ret := make([]int, k+1)
	for _, num := range nums {
		for i := k; i >= 0; i-- {
			v := dp[i][num] + 1
			if i > 0 {
				v = max(v, ret[i-1])
			}
			dp[i][v] = v
			ret[i] = max(ret[i], v)
		}
	}
	return ret[k]
}
