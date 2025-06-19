package main

import (
	"sort"
)

func maximumBeauty2(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] >= items[j][1]
		}
		return items[i][0] < items[j][0]
	})
	for i := 1; i < len(items); i++ {
		items[i][1] = max(items[i][1], items[i-1][1])
	}
	ret := make([]int, len(queries))
	for i, q := range queries {
		idx := sort.Search(len(items), func(i int) bool {
			return items[i][0] >= q
		})
		if idx == 0 && items[0][0] > q {
			continue
		}
		if idx == len(items) || items[idx][0] > q {
			ret[i] = items[idx-1][1]
			continue
		}
		ret[i] = items[idx][1]
	}
	return ret
}
