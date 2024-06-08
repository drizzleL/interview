package main

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	var sum int
	for _, ap := range apple {
		sum += ap
	}
	sort.Ints(capacity)
	var ret int
	for i := len(capacity) - 1; i >= 0 && sum > 0; i-- {
		sum -= capacity[i]
		ret += 1
	}
	return ret
}
