package main

import (
	"sort"
)

func splitPainting(segments [][]int) [][]int64 {
	start, end := map[int][][]int{}, map[int][][]int{}
	dict := map[int]bool{}
	var idxs []int
	for _, seg := range segments {
		start[seg[0]] = append(start[seg[0]], seg)
		if !dict[seg[0]] {
			idxs = append(idxs, seg[0])
			dict[seg[0]] = true
		}
		end[seg[1]] = append(end[seg[1]], seg)
		if !dict[seg[1]] {
			idxs = append(idxs, seg[1])
			dict[seg[1]] = true
		}
	}
	sort.Ints(idxs)
	var startIdx int
	var cnt int
	var currSum int
	var ret [][]int64
	for _, idx := range idxs {
		if cnt != 0 {
			ret = append(ret, []int64{int64(startIdx), int64(idx), int64(currSum)})
		}
		for _, seg := range start[idx] {
			cnt++
			currSum += seg[2]
		}
		for _, seg := range end[idx] {
			cnt--
			currSum -= seg[2]
		}
		startIdx = idx
	}
	return ret
}
