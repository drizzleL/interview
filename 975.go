package main

import (
	"sort"
)

func oddEvenJumps(arr []int) int {
	greater, less := make([]int, len(arr)), make([]int, len(arr))
	var q []int
	var pairs [][2]int
	for i := range arr {
		pairs = append(pairs, [2]int{i, arr[i]})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] < pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1]
	})
	for _, p := range pairs {
		i := p[0]
		greater[i] = i
		for len(q) != 0 && q[len(q)-1] < i {
			greater[q[len(q)-1]] = i
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	q = q[:0]
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] < pairs[j][0]
		}
		return pairs[i][1] > pairs[j][1]
	})
	for _, p := range pairs {
		i := p[0]
		less[i] = i
		for len(q) != 0 && q[len(q)-1] < i {
			less[q[len(q)-1]] = i
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	dict := make([][]int, len(arr))
	for i := range dict {
		dict[i] = make([]int, 2)
	}
	dict[len(arr)-1][0] = 1
	dict[len(arr)-1][1] = 1
	ret := 1
	for i := len(arr) - 2; i >= 0; i-- {
		if less[i] != i {
			dict[i][1] = dict[less[i]][0]
		}
		if greater[i] != i {
			dict[i][0] = dict[greater[i]][1]
		}
		ret += dict[i][0]
	}
	return ret
}
