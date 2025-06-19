package main

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	var ret int
	for _, v := range reward2 {
		ret += v
	}
	diff := make([]int, 0, len(reward1))
	for i := range reward1 {
		diff = append(diff, reward1[i]-reward2[i])
	}
	sort.Ints(diff)
	for i := 0; i < k; i++ {
		j := len(diff) - i - 1
		ret += diff[j]
	}
	return ret
}
