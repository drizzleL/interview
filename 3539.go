package main

import (
	"log"
	"math/bits"
)

func fact(x int) int {
	ret := 1
	for i := 1; i <= x; i++ {
		ret *= i
		ret %= 1e9 + 7
	}
	return ret
}

func magicalSum(m int, k int, nums []int) int {
	cache := map[[4]int]int{}
	var helper func(i int, size, uniq int, flag int) int
	helper = func(i int, m, k int, flag int) (ret int) {
		if m < 0 || k < 0 || m+bits.OnesCount(uint(flag)) < k {
			return 0
		}
		if m == 0 {
			if k == bits.OnesCount(uint(flag)) {
				return 1
			}
			return 0
		}
		if i == len(nums) {
			return 0
		}
		if c, ok := cache[[4]int{i, m, k, flag}]; ok {
			return c
		}
		defer func() {
			cache[[4]int{i, m, k, flag}] = ret
		}()
		for j := 0; j <= m; j++ {
			tmp := fastPow(nums[i], j, combination(m, j))
			newFlag := (flag + j) / 2
			newK := k - (flag+j)&1
			tmp *= helper(i+1, m-j, newK, newFlag)
			ret += tmp
			ret %= 1e9 + 7
		}
		return
	}
	return helper(0, m, k, 0)
}

func magicalSum2(m int, k int, nums []int) int {
	// pick k from [0,len(nums)-1]
	var arr2 []int
	for _, num := range nums {
		arr2 = append(arr2, num*num)
	}
	// pick pairs from arr2, can duplicate
	factor := 1
	pairs := (m - k) / 2
	if pairs != 0 {
		dp := make([][]int, len(arr2))
		for i := range dp {
			dp[i] = make([]int, pairs+1)
		}
		for i := 0; i < len(arr2); i++ {
			dp[i][0] = 1
		}
		for j := 1; j <= pairs; j++ {
			dp[0][j] = dp[0][j-1] * arr2[0]
		}
		for i := 1; i < len(arr2); i++ {
			for j := 1; j <= pairs; j++ {
				val := 1
				for m := 0; m <= j; m++ {
					dp[i][j] += dp[i-1][j-m] * val
					dp[i][j] %= 1e9 + 7
					val *= arr2[i]
					val %= 1e9 + 7
				}
			}
		}
		factor = dp[len(arr2)-1][pairs]
	}
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	for i := 0; i < len(nums); i++ {
		dp[i][0] = 1
	}
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]*nums[i]
			// dp[i][j] %= 1e9 + 7
		}
	}
	ret := dp[len(nums)-1][k] * factor
	log.Println(ret % (1e9 + 7))
	ret %= 1e9 + 7
	return ret
}
