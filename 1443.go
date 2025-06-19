package main

func minTime(n int, edges [][]int, hasApple []bool) int {
	dict := make([][]int, n)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	var dfs func(i int, p int) (bool, int)
	dfs = func(i int, p int) (bool, int) {
		has := hasApple[i]
		var ret int
		for _, next := range dict[i] {
			if next == p {
				continue
			}
			nextHas, tmp := dfs(next, i)
			if !nextHas {
				continue
			}
			has = true
			ret += tmp + 2
		}
		if !has {
			return false, 0
		}
		return true, ret
	}
	_, ret := dfs(0, -1)
	return ret
}
