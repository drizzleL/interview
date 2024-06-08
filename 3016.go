package main

import "sort"

func minimumPushes(word string) int {
	var items sort.IntSlice
	dict := map[byte]int{}
	for _, c := range word {
		if old, ok := dict[byte(c)]; ok {
			items[old] += 1
			continue
		}
		dict[byte(c)] = len(items)
		items = append(items, 1)
	}
	sort.Sort(sort.Reverse(items))
	var ret int
	cnt := 1
	for j := 0; j < len(items); {
		for i := 0; i < 8 && j < len(items); i, j = i+1, j+1 {
			ret += cnt * items[j]
		}
		cnt += 1
	}
	return ret
}
