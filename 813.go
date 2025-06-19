package main

func largestSumOfAverages(nums []int, k int) float64 {
	dp := make([][][]float64, len(nums))
	for i := range dp {
		dp[i] = make([][]float64, len(nums))
		for j := range dp[i] {
			dp[i][j] = make([]float64, k+1)
		}
	}
	sums := make([]int, len(nums)+1)
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	score := func(i, j int) float64 {
		sum := sums[j+1] - sums[i]
		return float64(sum) / float64(j-i+1)
	}
	var helper func(i, j, k int) float64
	helper = func(i, j, k int) (ret float64) {
		if dp[i][j][k] != 0 {
			return dp[i][j][k]
		}
		defer func() {
			dp[i][j][k] = ret
		}()
		if k == 1 {
			return score(i, j)
		}
		for mid := i; mid < j; mid++ {
			tmp := helper(i, mid, 1) + helper(mid+1, j, k-1)
			if tmp > ret {
				ret = tmp
			}
		}
		return
	}
	return helper(0, len(nums)-1, k)
}
