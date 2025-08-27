package main

import "math"

func numberOfWays2(n int, x int) int {
	dict := make([]int, n+1)
	dict[0] = 1
	for i := 1; ; i++ {
		num := int(math.Pow(float64(i), float64(x)))
		if num > n {
			break
		}
		for j := n; j >= num; j-- {
			dict[j] += dict[j-num]
			dict[j] %= 1e9 + 7
		}
	}
	return dict[n]
}
