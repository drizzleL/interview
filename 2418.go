package main

import "sort"

func sortPeople(names []string, heights []int) []string {
	type item struct {
		name   string
		height int
	}
	var items []*item
	for i := range names {
		items = append(items, &item{
			name:   names[i],
			height: heights[i],
		})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].height > items[j].height
	})
	var ret []string
	for _, item := range items {
		ret = append(ret, item.name)
	}
	return ret
}
