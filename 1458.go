package main

func maxDotProduct(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = nums1[i] * nums2[j]
			if i != 0 && j != 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+nums1[i]*nums2[j])
			}
			if i != 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
			}
			if j != 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

// func maxDotProduct(nums1 []int, nums2 []int) int {
// 	cache := map[[2]int]int{}
// 	var pick func(i, j int) (ret int)
// 	pick = func(i, j int) (ret int) {
// 		if c, ok := cache[[2]int{i, j}]; ok {
// 			return c
// 		}
// 		if i >= len(nums1) || j >= len(nums2) {
// 			return 0
// 		}
// 		if i == 0 && j == 0 {
// 			ret = math.MinInt32
// 		}
// 		defer func() {
// 			cache[[2]int{i, j}] = ret
// 		}()
// 		for i2 := i; i2 < len(nums1); i2++ {
// 			for j2 := j; j2 < len(nums2); j2++ {
// 				v := nums1[i2] * nums2[j2]
// 				if (i != 0 || j != 0) && v <= 0 {
// 					continue
// 				}
// 				ret = max(ret, v+pick(i2+1, j2+1))
// 			}
// 		}
// 		return ret
// 	}
// 	return pick(0, 0)
// }
