package main

func minZeroArray2(nums []int, queries [][]int) int {
	dp := make([][]bool, len(nums))
	var cnt int
	for i, num := range nums {
		if num == 0 {
			continue
		}
		dp[i] = make([]bool, num+1)
		dp[i][0] = true
		cnt += 1
	}
	if cnt == 0 {
		return 0
	}
	for i, q := range queries {
		for j := q[0]; j <= q[1]; j++ {
			if nums[j] == 0 {
				continue
			}
			for k := nums[j]; k-q[2] >= 0; k-- {
				dp[j][k] = dp[j][k] || dp[j][k-q[2]]
			}
			if dp[j][len(dp[j])-1] {
				nums[j] = 0
				cnt -= 1
			}
		}
		if cnt == 0 {
			return i + 1
		}
	}
	return -1
}
