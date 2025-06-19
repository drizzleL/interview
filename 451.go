package main

import "sort"

func frequencySort2(s string) string {
	dict := map[byte]int{}
	for _, c := range s {
		dict[byte(c)] += 1
	}
	type item struct {
		c    byte
		freq int
	}
	var items []*item
	for k, v := range dict {
		items = append(items, &item{
			c:    k,
			freq: v,
		})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].freq > items[j].freq
	})
	var b []byte
	for _, item := range items {
		for i := 0; i < item.freq; i++ {
			b = append(b, item.c)
		}
	}
	return string(b)
}
