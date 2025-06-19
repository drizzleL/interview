package main

func validArrangement(pairs [][]int) [][]int {
	dict := map[int][]int{}
	degree := map[int]int{}
	for _, p := range pairs {
		dict[p[0]] = append(dict[p[0]], p[1])
		degree[p[0]] += 1
		degree[p[1]] -= 1
	}
	start := pairs[0][0]
	for k, v := range degree {
		if v == 1 {
			start = k
			break
		}
	}
	nodes := make([]int, 0, len(pairs)+1)
	var dfs func(x int)
	dfs = func(x int) {
		for len(dict[x]) != 0 {
			y := dict[x][0]
			dict[x] = dict[x][1:]
			dfs(y)
		}
		nodes = append(nodes, x)
	}
	dfs(start)
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
	ret := make([][]int, 0, len(pairs))
	for i := 1; i < len(nodes); i++ {
		ret = append(ret, []int{nodes[i-1], nodes[i]})
	}
	return ret
}
