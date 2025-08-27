package main

import "sort"

func findMaximumLength(nums []int) int {
	q := [][3]int{{0, 0, 0}}
	var presum int
	for _, num := range nums {
		presum += num
		idx := sort.Search(len(q), func(i int) bool {
			return q[i][1] >= presum+1
		})
		q = q[:idx]
		top := q[len(q)-1]
		newVal := [3]int{presum, presum*2 - top[0], top[2] + 1}
		for q[len(q)-1][1] >= newVal[1] {
			q = q[:len(q)-1]
		}
		q = append(q, newVal)
	}
	return q[len(q)-1][2]
}
