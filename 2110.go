package main

func getDescentPeriods(prices []int) int64 {
	helper := func(x int) int {
		return x * (x + 1) / 2
	}
	var ret int
	for i := 0; i < len(prices); {
		j := i + 1
		for ; j < len(prices) && prices[j] == prices[j-1]-1; j++ {
		}
		ret += helper(j - i)
		i = j
	}
	return int64(ret)
}
