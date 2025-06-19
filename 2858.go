package main

func minEdgeReversals(n int, edges [][]int) []int {
	to := make([][]int, n)
	from := make([][]int, n)
	for _, ed := range edges {
		to[ed[0]] = append(to[ed[0]], ed[1])
		from[ed[1]] = append(from[ed[1]], ed[0])
	}
	seen1 := make([]bool, n)
	var dfs1 func(i int) int
	dfs1 = func(i int) int {
		seen1[i] = true
		var ret int
		for _, v := range to[i] {
			if seen1[v] {
				continue
			}
			ret += dfs1(v)
		}
		for _, v := range from[i] {
			if seen1[v] {
				continue
			}
			ret += 1 + dfs1(v)
		}
		return ret
	}
	ret := make([]int, n)
	seen2 := make([]bool, n)
	var dfs2 func(i int, pre int)
	dfs2 = func(i int, pre int) {
		ret[i] = pre
		seen2[i] = true
		for _, v := range to[i] {
			if seen2[v] {
				continue
			}
			dfs2(v, pre+1)
		}
		for _, v := range from[i] {
			if seen2[v] {
				continue
			}
			dfs2(v, pre-1)
		}
	}
	dfs2(0, dfs1(0))
	return ret
}
