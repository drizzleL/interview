package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num]++
	}
	tt := [][2]int{}
	for k, v := range dict {
		tt = append(tt, [2]int{k, v})
	}
	sort.Slice(tt, func(i, j int) bool {
		return tt[i][1] > tt[j][1]
	})
	var ret []int
	for i := 0; i < k; i++ {
		ret = append(ret, tt[i][0])
	}
	return ret
}
