package main

func maxOutput(n int, edges [][]int, price []int) int64 {
	dict := make([][]int, n)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	var ret int
	var dfs func(i int, pre int) (int, int)
	dfs = func(i int, pre int) (int, int) {
		maxPath, maxPath2 := price[i], 0
		for _, next := range dict[i] {
			if next == pre {
				continue
			}
			m1, m2 := dfs(next, i)
			ret = max(ret, max(maxPath2+m1, maxPath+m2))
			maxPath = max(maxPath, m1+price[i])
			maxPath2 = max(maxPath2, m2+price[i])
		}
		return maxPath, maxPath2
	}
	dfs(0, -1)
	return int64(ret)
}
