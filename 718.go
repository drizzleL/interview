package main

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	var ret int
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] != nums2[j] {
				continue
			}
			dp[i][j] = dp[i-1][j-1] + 1
			ret = max(ret, dp[i][j])
		}
	}
	return ret
}
