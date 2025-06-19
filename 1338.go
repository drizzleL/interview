package main

import "sort"

func minSetSize(arr []int) int {
	dict := map[int]int{}
	for _, num := range arr {
		dict[num]++
	}
	var freq []int
	for _, v := range dict {
		freq = append(freq, v)
	}
	sort.Ints(freq)
	var cnt int
	for i := len(freq) - 1; i >= 0; i-- {
		cnt += freq[i]
		if cnt >= len(arr)/2 {
			return len(arr) - i
		}
	}
	return 0
}
