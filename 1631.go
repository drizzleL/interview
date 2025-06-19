package main

import (
	"math"
)

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	if m == 1 && n == 1 {
		return 0
	}
	l, r := math.MaxInt32, math.MinInt32
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != 0 {
				tmp := abs(heights[i-1][j] - heights[i][j])
				l = min(l, tmp)
				r = max(r, tmp)
			}
			if j != 0 {
				tmp := abs(heights[i][j-1] - heights[i][j])
				l = min(l, tmp)
				r = max(r, tmp)
			}
		}
	}
	toXy := func(num int) (int, int) {
		return num / n, num % n
	}
	toNum := func(x, y int) int {
		return x*n + y
	}
	check := func(effort int) bool {
		seen := make([]bool, m*n)
		nodes := []int{0}
		for len(nodes) > 0 {
			var next []int
			for _, node := range nodes {
				seen[node] = true
				x, y := toXy(node)
				if x == m-1 && y == n-1 {
					return true
				}
				for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
					x2, y2 := x+dir[0], y+dir[1]
					if x2 < 0 || y2 < 0 || x2 >= m || y2 >= n {
						continue
					}
					node2 := toNum(x2, y2)
					if seen[node2] {
						continue
					}
					if abs(heights[x][y]-heights[x2][y2]) > effort {
						continue
					}
					seen[node2] = true
					next = append(next, node2)
				}
			}
			nodes = next
		}
		return false
	}
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1

		}
	}
	return l
}
