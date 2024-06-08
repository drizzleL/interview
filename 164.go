package main

import (
	"math"
)

func maximumGap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	m, M := math.MaxInt32, 0
	used := map[int]bool{}
	for _, num := range nums { // find buckets
		m = min(m, num)
		M = max(M, num)
		used[num] = true
	}
	if m == M {
		return 0
	}
	gap := (M - m) / (len(used) - 1)
	buckets := make([][2]int, len(used))
	for i := range buckets {
		buckets[i][0] = math.MaxInt32
		buckets[i][1] = -1
	}
	for num := range used {
		n := (num - m) / (gap + 1)
		buckets[n][0] = min(buckets[n][0], num)
		buckets[n][1] = max(buckets[n][1], num)
	}
	var maxGap int
	start := buckets[0][0]
	for i := 0; i < len(buckets); i++ {
		if buckets[i][0] != math.MaxInt32 {
			maxGap = max(maxGap, buckets[i][0]-start)
			start = buckets[i][0]
		}
		if buckets[i][1] >= 0 {
			maxGap = max(maxGap, buckets[i][1]-start)
			start = buckets[i][1]
		}
	}
	return maxGap
}
