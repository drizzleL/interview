package main

import "sort"

func minimumOperations8(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	dict := [2]map[int]int{}
	for i := range dict {
		dict[i] = make(map[int]int)
	}
	for i, num := range nums {
		dict[i%2][num] += 1
	}
	largest := func(m map[int]int) [][2]int {
		var items [][2]int
		for k, v := range m {
			items = append(items, [2]int{k, v})
		}
		sort.Slice(items, func(i, j int) bool {
			return items[i][1] > items[j][1]
		})
		if len(items) == 1 {
			return items
		}
		return items[:2]
	}
	a, b := largest(dict[0]), largest(dict[1])
	if a[0][0] != b[0][0] {
		return len(nums) - a[0][1] - b[0][1]
	}
	if len(a) == 1 && len(b) == 1 {
		return len(nums) / 2
	}
	if len(a) == 1 {
		return len(nums) - max(a[0][1]+b[1][1], b[0][1])
	}
	if len(b) == 1 {
		return len(nums) - max(a[0][1], a[1][1]+b[0][1])
	}
	return len(nums) - max(a[0][1]+b[1][1], a[1][1]+b[0][1])
}
