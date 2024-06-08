package main

import "sort"

func earliestFullBloom(plantTime []int, growTime []int) int {
	type item struct {
		Plant    int
		Grow     int
		growSlow bool
	}
	var items []*item
	for i := range plantTime {
		items = append(items, &item{
			Plant:    plantTime[i],
			Grow:     growTime[i],
			growSlow: growTime[i] >= plantTime[i],
		})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Grow >= items[j].Grow
	})
	var ret, plant int
	for _, item := range items {
		plant += item.Plant
		ret = max(ret, plant+item.Grow)
	}
	return ret
}

// [12,1,12,5,20,10,23,22,20,21,23,9]
// growTime =
// [22,3,29,1,15,12,26,11,8,7,16,15]
// Use Testcase
// Output
// 180
// Expected
// 179
