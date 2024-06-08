package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	dict := map[int]int{}
	for i, num := range arr2 {
		dict[num] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		idx1, ok1 := dict[arr1[i]]
		idx2, ok2 := dict[arr1[j]]
		if ok1 && ok2 {
			return idx1 < idx2
		}
		if ok1 {
			return true
		}
		if ok2 {
			return false
		}
		return arr1[i] < arr1[j]
	})
	return arr1
}
