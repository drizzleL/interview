package main

import "log"

func maxProduct2(nums []int, k int, limit int) int {
	var sum, zeroCnt int
	for _, num := range nums {
		sum += num
		if num == zeroCnt {
			zeroCnt += 1
		}
	}
	dp := make([][][]bool, 3)
	for i := range dp {
		dp[i] = make([][]bool, limit+1)
		for j := range dp[i] {
			dp[i][j] = make([]bool, sum*2+1)
		}
	}
	if k+sum >= sum*2+1 || k+sum < 0 {
		return -1
	}
	sumDp := make([][]bool, 3)
	for i := range dp {
		sumDp[i] = make([]bool, sum*2+1)
	}
	ret := -1
	for i := 0; i < len(nums); i++ {
		for preSum := len(sumDp[0]) - 1; preSum >= 0; preSum-- {
			if sumDp[0][preSum] {
				sumDp[1][preSum-nums[i]] = true
			}
			if sumDp[1][preSum] {
				sumDp[2][preSum+nums[i]] = true
			}
			if sumDp[2][preSum] {
				sumDp[1][preSum-nums[i]] = true
			}
		}
		sumDp[0][nums[i]+sum] = true
		if zeroCnt >= 1 {
			for j := 0; j <= 2; j++ {
				if sumDp[j][k+sum] {
					log.Println(111)
					ret = max(ret, 0)
				}
			}
		}
		for preProd := len(dp[0]) - 1; preProd >= 0; preProd-- {
			if preProd*nums[i] > limit {
				continue
			}
			if nums[i] == 0 {
				for preSum := len(dp[0][preProd]) - 1; preSum >= 0; preSum-- {
					if dp[1][preProd][preSum] {
						dp[2][0][preSum] = true
					}
					if dp[2][preProd][preSum] {
						dp[1][0][preSum] = true
					}
					if dp[0][preProd][preSum] {
						dp[1][0][preSum] = true
					}
				}
			} else {
				for preSum := len(dp[0][preProd]) - 1; preSum >= 0; preSum-- {
					if dp[0][preProd][preSum] {
						dp[1][preProd*nums[i]][preSum-nums[i]] = true
					}
					if dp[1][preProd][preSum] {
						dp[2][preProd*nums[i]][preSum+nums[i]] = true
					}
					if dp[2][preProd][preSum] {
						dp[1][preProd*nums[i]][preSum-nums[i]] = true
					}
				}
			}
		}
		if nums[i] <= limit {
			dp[0][nums[i]][nums[i]+sum] = true
		}
		if nums[i] == 0 {
			zeroCnt -= 1
		}
	}
	for _, v := range dp {
		for prod := range v {
			if !v[prod][k+sum] {
				continue
			}
			ret = max(ret, prod)
		}
	}
	return ret
}
