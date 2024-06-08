package main

func wardrobeFinishing(m int, n int, cnt int) int {
	getDigit := func(v int) int {
		var ret int
		for v != 0 {
			ret += v % 10
			v /= 10
		}
		return ret
	}
	visited := make([]bool, m*n)
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= m || j >= n || visited[i*n+j] || getDigit(i)+getDigit(j) > cnt {
			return 0
		}
		visited[i*n+j] = true
		return 1 + dfs(i+1, j) + dfs(i, j+1)
	}
	return dfs(0, 0)
}
