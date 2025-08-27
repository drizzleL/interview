package main

func maximumLength6(nums []int, k int) int {
	arr := make([][]int, len(nums))
	for i := range arr {
		arr[i] = make([]int, k+1)
	}
	lastDict := map[int]int{}
	dp := make([]int, k+1)
	for i, num := range nums {
		arr[i][0] = 1
		for j := k; j >= 1; j-- {
			arr[i][j] = dp[j-1] + 1
		}
		if c, ok := lastDict[num]; ok {
			for j := 0; j <= k; j++ {
				arr[i][j] = max(arr[i][j], arr[c][j]+1)
			}
		}
		lastDict[num] = i
		for j := 0; j <= k; j++ {
			dp[j] = max(dp[j], arr[i][j])
		}
	}
	return dp[k]
}
