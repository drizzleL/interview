package main

import "sort"

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	var ret []int
	check := func(start, end int) int {
		i := sort.SearchInts(ret, start)
		j := sort.SearchInts(ret, end+1)
		return j - i
	}
	for _, intv := range intervals {
		start, end := intv[0], intv[1]
		has := check(start, end)
		if has >= 2 {
			continue
		}
		if has == 0 {
			ret = append(ret, end-1, end)
			continue
		}
		if ret[len(ret)-1] == end {
			ret = ret[:len(ret)-1]
			ret = append(ret, end-1, end)
			continue
		}
		ret = append(ret, end)
	}
	return len(ret)
}
