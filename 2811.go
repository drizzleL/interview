package main

func canSplitArray(nums []int, m int) bool {
	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, len(nums))
		dp[i][i] = true
		if i+1 < len(nums) {
			dp[i][i+1] = true
		}
	}
	sums := make([]int, len(nums)+1)
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for size := 3; size <= len(nums); size++ {
		for i := 0; i+size-1 < len(nums); i++ {
			j := i + size - 1
			for mid := i; mid < j; mid++ {
				if !dp[i][mid] {
					continue
				}
				if !dp[mid+1][j] {
					continue
				}
				if mid-i+1 != 1 && sums[mid+1]-sums[i] < m {
					continue
				}
				if j-mid != 1 && sums[j+1]-sums[mid+1] < m {
					continue
				}
				dp[i][j] = true
				break
			}
		}
	}
	return dp[0][len(nums)-1]
}
