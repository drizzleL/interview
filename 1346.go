package main

import "sort"

func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	dict := map[int]bool{}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] >= 0 {
			if dict[arr[i]*2] {
				return true
			}
		} else {
			if arr[i]%2 == 0 && dict[arr[i]/2] {
				return true
			}
		}
		dict[arr[i]] = true
	}
	return false
}
