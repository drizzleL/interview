package main

import (
	"math"
	"sort"
)

func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)
	sort.Ints(walls)
	var eles [][2]int
	for i := range robots {
		eles = append(eles, [2]int{robots[i], distance[i]})
	}
	eles = append(eles, [2]int{math.MaxInt32, 0})
	sort.Slice(eles, func(i, j int) bool {
		return eles[i][0] < eles[j][0]
	})
	countWalls := func(l, r int) (ret int) {
		i := sort.SearchInts(walls, l)
		j := sort.SearchInts(walls, r+1)
		return j - i
	}
	dp := make([][2]int, n)
	dp[0][0] = countWalls(eles[0][0]-eles[0][1], eles[0][0])
	dp[0][1] = countWalls(eles[0][0], min(eles[0][0]+eles[0][1], eles[1][0]-1))
	for i := 1; i < n; i++ {
		right := countWalls(eles[i][0], min(eles[i][0]+eles[i][1], eles[i+1][0]-1))
		left1 := countWalls(max(eles[i-1][0]+1, eles[i][0]-eles[i][1]), eles[i][0])
		left2 := countWalls(max(eles[i-1][0]+eles[i-1][1]+1, eles[i][0]-eles[i][1]), eles[i][0])
		dp[i][0] = max(dp[i-1][0]+left1, dp[i-1][1]+left2)
		dp[i][1] = max(dp[i-1][0], dp[i-1][1]) + right
	}
	return max(dp[n-1][0], dp[n-1][1])
}
