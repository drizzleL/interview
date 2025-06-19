package main

import (
	"log"
	"math"
)

func minimumCoins(prices []int) int {
	dp := make([]int, len(prices)+1)
	for i := len(prices) - 1; i >= 0; i-- {
		dp[i] = math.MaxInt32
		for j := i + 1; j <= i+i && j <= len(prices); j++ {
			dp[i] = min(dp[i], prices[i]+dp[j])
		}
	}
	log.Println(dp)
	return dp[0]
}
