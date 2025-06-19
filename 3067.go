package main

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	dict := map[int][][2]int{}
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], [2]int{ed[1], ed[2] % signalSpeed})
		dict[ed[1]] = append(dict[ed[1]], [2]int{ed[0], ed[2] % signalSpeed})
	}
	var dfs func(i int, weight int, seen []bool) int
	dfs = func(i int, weight int, seen []bool) (ret int) {
		seen[i] = true
		if weight%signalSpeed == 0 {
			ret += 1
		}
		for _, child := range dict[i] {
			if seen[child[0]] {
				continue
			}
			seen[child[0]] = true
			ret += dfs(child[0], weight+child[1], seen)
		}
		return
	}
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		var cnt int
		seen := make([]bool, n)
		seen[i] = true
		for _, child := range dict[i] { // traverse every branch
			cnt2 := dfs(child[0], child[1], seen)
			ret[i] += cnt * cnt2
			cnt += cnt2
		}
	}
	return ret
}
