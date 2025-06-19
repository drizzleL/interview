package main

import "container/heap"

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	q := HeapArr{
		LessHelper: func(a, b interface{}) bool {
			return a.([3]int)[2] < b.([3]int)[2]
		},
	}
	for i := 0; i < m; i++ {
		visited[i][0] = true
		visited[i][n-1] = true
		heap.Push(&q, [3]int{i, 0, heightMap[i][0]})
		heap.Push(&q, [3]int{i, 0, heightMap[i][n-1]})
	}
	for j := 0; j < n; j++ {
		visited[0][j] = true
		visited[m-1][j] = true
		heap.Push(&q, [3]int{0, j, heightMap[0][j]})
		heap.Push(&q, [3]int{m - 1, j, heightMap[m-1][j]})
	}
	var ret int
	for q.Len() > 0 {
		top := heap.Pop(&q).([3]int)
		i, j, h := top[0], top[1], top[2]
		for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			i2, j2 := i+dir[0], j+dir[1]
			if i2 < 0 || j2 < 0 || i2 >= m || j2 >= n {
				continue
			}
			if visited[i2][j2] {
				continue
			}
			ret += max(0, h-heightMap[i2][j2])
			visited[i2][j2] = true
			heap.Push(&q, [3]int{i2, j2, max(h, heightMap[i2][j2])})
		}
	}
	return ret
}
