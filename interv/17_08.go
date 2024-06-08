package main

import (
	"sort"
)

func bestSeqAtIndex(height []int, weight []int) int {
	p := [][2]int{}
	for i := range height {
		h, w := height[i], weight[i]
		p = append(p, [2]int{h, w})
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i][0] == p[j][0] {
			return p[i][1] > p[j][1]
		}
		return p[i][0] < p[j][0]
	})
	dp := []int{}
	for i := range p {
		w := p[i][1]
		idx := sort.SearchInts(dp, w)
		if idx == len(dp) {
			dp = append(dp, w)
		} else {
			dp[idx] = w
		}
	}
	return len(dp)
}

func bestSeqAtIndex2(height []int, weight []int) int {
	p := [][2]int{}
	for i := range height {
		h, w := height[i], weight[i]
		p = append(p, [2]int{h, w})
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i][0] == p[j][0] {
			return p[i][1] <= p[j][1]
		}
		return p[i][0] < p[j][0]
	})
	cache := make([]int, len(p))
	for i := len(p) - 1; i >= 0; i-- {
		var v int
		for j := i + 1; j < len(p); j++ {
			if p[j][0] > p[i][0] && p[j][1] > p[i][1] {
				v = max(v, cache[j])
			}
		}
		cache[i] = v + 1
	}
	var ret int
	for _, v := range cache {
		ret = max(ret, v)
	}
	return ret
}
