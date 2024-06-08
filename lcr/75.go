package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	dict := map[int]int{}
	for i, v := range arr2 {
		dict[v] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		rank1, found1 := dict[arr1[i]]
		rank2, found2 := dict[arr1[j]]
		if found1 && found2 {
			return rank1 < rank2
		}
		if !found1 && !found2 {
			return arr1[i] < arr1[j]
		}
		if found1 {
			return true
		}
		return false
	})
	return arr1
}
