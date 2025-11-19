package main

import (
	"container/list"
	"math"
)

func minimumVisitedCells(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	dist[0][0] = 1
	queue := list.New()
	queue.PushBack([3]int{0, 0, 1})

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		curr := element.Value.([3]int)
		r, c, d := curr[0], curr[1], curr[2]
		if r == m-1 && c == n-1 {
			return d
		}
		// Move right
		for nc := c + 1; nc <= c+grid[r][c] && nc < n; nc++ {
			if dist[r][nc] > d+1 {
				dist[r][nc] = d + 1
				queue.PushBack([3]int{r, nc, d + 1})
			}
		}
		// Move down
		for nr := r + 1; nr <= r+grid[r][c] && nr < m; nr++ {
			if dist[nr][c] > d+1 {
				dist[nr][c] = d + 1
				queue.PushBack([3]int{nr, c, d + 1})
			}
		}
	}

	return -1
}
