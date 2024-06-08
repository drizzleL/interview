package main

func allPathsSourceTarget(graph [][]int) [][]int {
	visited := make([]bool, len(graph))
	var helper func(x int, curr []int)
	var ret [][]int
	helper = func(x int, curr []int) {
		if x == len(graph)-1 {
			curr = append(curr, x)
			ret = append(ret, append([]int{}, curr...))
			return
		}
		if visited[x] {
			return
		}
		visited[x] = true
		for i := 0; i < len(graph[x]); i++ {
			helper(graph[x][i], append(curr, x))
		}
		visited[x] = false
	}
	helper(0, nil)
	return ret
}
