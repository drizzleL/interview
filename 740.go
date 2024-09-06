package main

import (
	"sort"
)

func deleteAndEarn(nums []int) int {
	var newNums []int
	dict := map[int]int{}
	for _, num := range nums {
		dict[num] += 1
		if dict[num] == 1 {
			newNums = append(newNums, num)
		}
	}
	sort.Ints(newNums)
	var a, b int
	b = newNums[0] * dict[newNums[0]]
	for i := 1; i < len(newNums); i++ {
		if newNums[i] == newNums[i-1]+1 {
			a, b = max(a, b), a+newNums[i]*dict[newNums[i]]
		} else {
			a, b = max(a, b), max(a, b)+newNums[i]*dict[newNums[i]]
		}
	}
	return max(a, b)
}
