package main

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	in := make([]int, n)
	dict := make([][]int, n)
	for _, ed := range edges {
		in[ed[1]] += 1
		dict[ed[0]] = append(dict[ed[0]], ed[1])
	}
	var nodes []int
	nodeToColors := make([][26]int, n)
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			nodes = append(nodes, i)
		}
	}
	ret := make([]int, len(colors))
	for i := range ret {
		ret[i] = -1
	}
	for len(nodes) != 0 {
		node := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		idx := int(colors[node] - 'a')
		c := nodeToColors[node]
		c[idx] += 1
		ret[node] = max(ret[node], c[idx])
		for _, next := range dict[node] {
			for i := range c {
				nodeToColors[next][i] = max(nodeToColors[next][i], c[i])
			}
			in[next] -= 1
			if in[next] == 0 {
				nodes = append(nodes, next)
			}
		}
	}
	var ans int
	for _, v := range ret {
		if v == -1 {
			return -1
		}
		ans = max(ans, v)
	}
	return ans
}
