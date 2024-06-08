package main

import (
	"sort"
)

func sortJumbled(mapping []int, nums []int) []int {
	type item struct {
		val int
		idx int
	}
	var items []*item
	for i, num := range nums {
		it := &item{
			idx: i,
		}
		items = append(items, it)
		if num == 0 {
			it.val = mapping[0]
		}
		base := 1
		for num != 0 {
			it.val += base * mapping[num%10]
			num /= 10
			base *= 10
		}
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].val == items[j].val {
			return items[i].idx < items[j].idx
		}
		return items[i].val < items[j].val
	})
	var ret []int
	for _, item := range items {
		ret = append(ret, nums[item.idx])
	}
	return ret
}
