package main

import "math"

func minCost20(n int, cost [][]int) int64 {
	pre := make([]int, 9)
	for i, j := n/2-1, n/2; i >= 0; i, j = i-1, j+1 {
		next := make([]int, 9)
		for i := range next {
			next[i] = math.MaxInt64
		}
		for a := 0; a < 3; a++ {
			for b := 0; b < 3; b++ {
				if a == b {
					continue
				}
				c1, c2 := cost[i][a], cost[j][b]
				if i == n/2-1 {
					next[a*3+b] = c1 + c2
					continue
				}
				for k, v := range pre {
					if v == math.MaxInt64 {
						continue
					}
					a1, b1 := k/3, k%3
					if a1 == a || b1 == b {
						continue
					}
					next[a*3+b] = min(next[a*3+b], v+c1+c2)
				}
			}
		}
		pre = next
	}
	ret := math.MaxInt64
	for _, v := range pre {
		ret = min(ret, v)
	}
	return int64(ret)
}
