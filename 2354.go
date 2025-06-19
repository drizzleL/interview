package main

import (
	"sort"
)

func countExcellentPairs(nums []int, k int) int64 {
	dict := map[int]int{}
	for _, num := range nums {
		if dict[num] != 0 {
			continue
		}
		var cnt int
		for num2 := num; num2 != 0; num2 >>= 1 {
			if num2%2 == 1 {
				cnt += 1
			}
		}
		dict[num] = cnt
	}
	var ret int
	var cnts []int
	for _, v := range dict {
		cnts = append(cnts, v)
	}
	sort.Ints(cnts)
	for i := 0; i < len(cnts); i++ {
		idx := sort.SearchInts(cnts, k-cnts[i])
		ret += len(cnts) - idx
	}
	return int64(ret)
}
