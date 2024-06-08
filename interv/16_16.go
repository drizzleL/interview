package main

import (
	"sort"
)

func subSort(array []int) []int {
	if len(array) == 0 {
		return []int{-1, -1}
	}
	var m int
	for ; m < len(array)-1; m++ {
		if array[m+1] < array[m] {
			break
		}
	}
	if m == len(array)-1 {
		return []int{-1, -1}
	}
	n := m + 1
	i := n
	maxVal := array[m]
	for ; i < len(array); i++ {
		if array[i] >= maxVal {
			maxVal = array[i]
			continue
		}
		n = i
		m = sort.Search(m, func(j int) bool {
			return array[j] > array[i]
		})
	}
	return []int{m, n}
}
