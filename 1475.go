package main

func finalPrices(prices []int) []int {
	var q []int
	ret := make([]int, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {
		for len(q) != 0 && q[len(q)-1] > prices[i] {
			q = q[:len(q)-1]
		}
		ret[i] = prices[i]
		if len(q) != 0 {
			ret[i] -= q[len(q)-1]
		}
		q = append(q, prices[i])
	}
	return ret
}
