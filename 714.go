package main

import "math"

func maxProfit(prices []int, fee int) int {
	buy := math.MinInt32
	sell := 0
	for _, p := range prices {
		sell, buy = max(sell, buy+p-fee), max(buy, sell-p)
	}
	return sell
}
