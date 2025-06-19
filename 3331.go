package main

func findSubtreeSizes(parent []int, s string) []int {
	children := make([][]int, len(parent))
	for i, p := range parent {
		if p == -1 {
			continue
		}
		children[p] = append(children[p], i)
	}
	seen := [26]int{}
	for i := range seen {
		seen[i] = -1
	}
	var dfs func(x int)
	dfs = func(x int) {
		v := s[x] - 'a'
		old := seen[v]
		if old != -1 { // seen before
			parent[x] = old
		}
		seen[v] = x
		for _, child := range children[x] {
			dfs(child)
		}
		seen[v] = old
	}
	dfs(0)
	ret := make([]int, len(parent))
	children = make([][]int, len(parent))
	for i, p := range parent {
		if p == -1 {
			continue
		}
		children[p] = append(children[p], i)
	}
	var helper func(x int) int
	helper = func(x int) int {
		ret[x] = 1
		for _, child := range children[x] {
			ret[x] += helper(child)
		}
		return ret[x]
	}
	helper(0)
	return ret
}
