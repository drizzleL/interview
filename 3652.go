package main

func maxProfit8(prices []int, strategy []int, k int) int64 {
	var sum int
	presums := make([]int, len(prices)+1)
	absPresums := make([]int, len(prices)+1)
	for i, p := range prices {
		sum += p * strategy[i]
		presums[i+1] = presums[i] + p*strategy[i]
		absPresums[i+1] = absPresums[i] + p
	}
	getSpan := func(i, j int) int {
		return presums[j+1] - presums[i]
	}
	getAbsSpan := func(i, j int) int {
		return absPresums[j+1] - absPresums[i]
	}
	ret := sum
	for m := k - 1; m < len(prices); m++ { // end with m
		tmp := sum - getSpan(m-k+1, m) + getAbsSpan(m-k/2+1, m)
		ret = max(ret, tmp)
	}
	return int64(ret)
}
