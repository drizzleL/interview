package main

import "sort"

func smallestAbsent(nums []int) int {
	var sum int
	sort.Ints(nums)
	dict := map[int]bool{}
	for _, num := range nums {
		sum += num
		dict[num] = true
	}
	avg := sum / len(nums)
	for i := avg + 1; ; i++ {
		if !dict[i] {
			return i
		}
	}
}
