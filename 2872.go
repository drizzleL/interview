package main

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	dict := map[int][]int{}
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	var ret int
	seen := make([]bool, n)
	var dfs func(x int) int
	dfs = func(x int) int {
		seen[x] = true
		val := values[x]
		for _, child := range dict[x] {
			if seen[child] {
				continue
			}
			seen[child] = true
			val += dfs(child)
		}
		if val%k == 0 {
			ret += 1
			return 0
		}
		return val
	}
	dfs(0)
	return ret
}
