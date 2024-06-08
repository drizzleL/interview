package main

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	in := make([]int, n)
	for _, r := range roads {
		in[r[0]] += 1
		in[r[1]] += 1
	}
	var items []int
	for _, v := range in {
		items = append(items, v)
	}
	sort.Ints(items)
	var ret int
	for i := 0; i < len(items); i++ {
		ret += (i + 1) * items[i]
	}
	return int64(ret)
}
