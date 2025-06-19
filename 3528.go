package main

func baseUnitConversions(conversions [][]int) []int {
	n := len(conversions) + 1
	dict := make([][][2]int, n)
	for _, conv := range conversions {
		dict[conv[0]] = append(dict[conv[0]], [2]int{conv[1], conv[2]})
		dict[conv[1]] = append(dict[conv[1]], [2]int{conv[0], conv[2]})
	}
	seen := make([]bool, n)
	ret := make([]int, n)
	var dfs func(i, val int)
	dfs = func(i, val int) {
		val %= 1e9 + 7
		seen[i] = true
		ret[i] = val
		for _, next := range dict[i] {
			if seen[next[0]] {
				continue
			}
			dfs(next[0], val*next[1])
		}
	}
	dfs(0, 1)
	return ret
}
