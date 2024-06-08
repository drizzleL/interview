package main

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	if len(grid2) == 0 {
		return 0
	}
	m, n := len(grid2), len(grid2[0])
	visited := make([]bool, m*n)
	var dfs func(i, j int, tmp *[][]int)
	dfs = func(i, j int, tmp *[][]int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if grid2[i][j] == 0 || visited[i*n+j] {
			return
		}
		visited[i*n+j] = true
		*tmp = append(*tmp, []int{i, j})
		dfs(i-1, j, tmp)
		dfs(i+1, j, tmp)
		dfs(i, j+1, tmp)
		dfs(i, j-1, tmp)
	}
	var islands [][][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i*n+j] || grid2[i][j] == 0 {
				continue
			}
			var tmp [][]int
			dfs(i, j, &tmp)
			islands = append(islands, tmp)
		}
	}
	var ret int
	for _, isl := range islands {
		flag := 1
		for _, node := range isl {
			if grid1[node[0]][node[1]] == 0 {
				flag = 0
				continue
			}
		}
		ret += flag
	}
	return ret
}
