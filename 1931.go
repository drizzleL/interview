package main

import (
	"math"
)

func colorTheGrid(m int, n int) int {
	toVals := func(x int) []int {
		vals := make([]int, m)
		for i := 0; i < m; i++ {
			vals[i] = x % 3
			x /= 3
		}
		return vals
	}
	dict := map[int][]int{}
	maxVal := int(math.Pow(3, float64(m)))
	pre := make([]int, maxVal)
	check := func(vals []int) bool {
		for i := 1; i < len(vals); i++ {
			if vals[i] == vals[i-1] {
				return false
			}
		}
		return true
	}
	check2 := func(a, b []int) bool {
		for i := range a {
			if a[i] == b[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < maxVal; i++ {
		vals := toVals(i)
		if !check(vals) {
			continue
		}
		pre[i] = 1
		for j := 0; j < maxVal; j++ {
			vals2 := toVals(j)
			if !check(vals2) {
				continue
			}
			if !check2(vals, vals2) {
				continue
			}
			dict[i] = append(dict[i], j)
		}
	}
	for i := 1; i < n; i++ {
		next := make([]int, maxVal)
		for k, from := range dict {
			for _, m := range from {
				next[k] += pre[m]
				next[k] %= 1e9 + 7
			}
		}
		pre = next
	}
	var ret int
	for _, v := range pre {
		ret += v
		ret %= 1e9 + 7
	}
	return ret
}
