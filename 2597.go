package main

import "sort"

func beautifulSubsets(nums []int, k int) int {
	sort.Ints(nums)
	dict := map[int][]int{}
	for _, num := range nums {
		dict[num%k] = append(dict[num%k], num)
	}
	ret := 1
	for _, arr := range dict {
		tmp := 1
		cntDict := map[int]int{}
		for _, num := range arr {
			cnt := tmp - cntDict[num-k]
			tmp += cnt
			cntDict[num] += cnt
		}
		ret *= tmp
	}
	return ret - 1
}
