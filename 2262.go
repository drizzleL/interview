package main

import "sort"

func appealSum(s string) int64 {
	var dict [26][]int
	for i, c := range s {
		dict[c-'a'] = append(dict[c-'a'], i)
	}
	var ret int
	for i := range s {
		// c := int(s[i] - 'a')
		var idxs []int
		for _, v := range dict {
			nextIdx := sort.SearchInts(v, i)
			if nextIdx == len(v) {
				continue
			}
			idxs = append(idxs, v[nextIdx])
		}
		idxs = append(idxs, len(s))
		sort.Ints(idxs)
		for j := 1; j < len(idxs); j++ {
			ret += (idxs[j] - idxs[j-1]) * j
		}
	}
	return int64(ret)
}
