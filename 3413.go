package main

import (
	"sort"
)

func maximumCoins(coins [][]int, k int) int64 {
	dict := map[int]int{}
	startDict := map[int]bool{}
	for _, coin := range coins {
		dict[coin[0]] += coin[2]
		dict[coin[1]+1] -= coin[2]
	}
	var pos [][2]int
	for k, v := range dict {
		pos = append(pos, [2]int{k, v})
	}
	sort.Slice(pos, func(i, j int) bool {
		return pos[i][0] < pos[j][0]
	})
	starts := []int{}
	preSum := make([][3]int, len(pos)+1) // 0:idx 1:delta 2:presum
	for i, p := range pos {
		preSum[i+1] = [3]int{p[0], preSum[i][1] + p[1], preSum[i][2] + preSum[i][1]*(p[0]-preSum[i][0])}
	}
	for _, coin := range coins {
		startDict[coin[0]] = true
		if coin[1]-k+1 < pos[0][0] {
			continue
		}
		startDict[coin[1]-k+1] = true
	}
	for k := range startDict {
		starts = append(starts, k)
	}
	getPreSum := func(a int) int {
		idx := sort.Search(len(preSum), func(i int) bool {
			return preSum[i][0] > a
		}) - 1
		return preSum[idx][2] + (preSum[idx][1] * (a - preSum[idx][0] + 1))
	}
	getSpanSum := func(a, b int) int {
		return getPreSum(b) - getPreSum(a-1)
	}
	var ret int
	for _, start := range starts {
		ret = max(ret, getSpanSum(start, start+k-1))
	}
	return int64(ret)
}
