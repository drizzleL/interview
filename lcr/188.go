package main

func bestTiming(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	mPrice := prices[0]
	var ret int
	for i := 1; i < len(prices); i++ {
		if prices[i]-mPrice > ret {
			ret = prices[i] - mPrice
		}
		if prices[i] < mPrice {
			mPrice = prices[i]
		}
	}
	return ret
}
