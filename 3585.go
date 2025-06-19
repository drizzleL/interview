package main

import (
	"math"
)

func findMedian(n int, edges [][]int, queries [][]int) []int {
	dict := make([][][2]int, n)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], [2]int{ed[1], ed[2]})
		dict[ed[1]] = append(dict[ed[1]], [2]int{ed[0], ed[2]})
	}
	depth := make([]int, n)
	weight := make([]int, n)
	dirParent := make([]int, n)
	var dfs func(i, p, dep, w int) int
	dfs = func(i, p, dep, w int) int {
		dirParent[i] = p
		depth[i] = dep
		weight[i] = w
		var ret int
		for _, ed := range dict[i] {
			if ed[0] == p {
				continue
			}
			ret = max(ret, dfs(ed[0], i, dep+1, w+ed[1]))
		}
		return ret + 1
	}
	maxDep := dfs(0, -1, 0, 0)
	bitLen := int(math.Ceil(math.Log2(float64(maxDep)))) + 1
	parent := make([][]int, n)
	for i := range parent {
		parent[i] = make([]int, bitLen)
		for j := 0; j < int(bitLen); j++ {
			parent[i][j] = -1
		}
	}
	for i := 0; i < n; i++ {
		parent[i][0] = dirParent[i]
	}
	for j := 1; j < int(bitLen); j++ {
		for i := 0; i < n; i++ {
			if parent[i][j-1] == -1 {
				continue
			}
			parent[i][j] = parent[parent[i][j-1]][j-1]
		}
	}
	getLca := func(a, b int) int {
		if depth[a] > depth[b] {
			a, b = b, a
		}
		d := depth[b] - depth[a]
		for d > 0 { //
			step := int(math.Log2(float64(d)))
			b = parent[b][step]
			d -= 1 << step
		}
		if a == b {
			return a
		}
		for d := bitLen - 1; d >= 0; d-- {
			if parent[a][d] != -1 && parent[a][d] != parent[b][d] {
				a, b = parent[a][d], parent[b][d]
			}
		}
		return parent[a][0]
	}
	ret := make([]int, len(queries))
	getKthAncestor := func(x int, k int) int {
		for m := bitLen - 1; x != -1 && m >= 0; m-- {
			if (1<<m)&k != 0 {
				x = parent[x][m]
			}
		}
		return x
	}
	findBetween := func(u, lca int, halfCost float64, less int) int {
		low, high := 0, depth[u]-depth[lca]
		ans := u
		for low <= high {
			mid := (low + high) / 2
			x := getKthAncestor(u, mid)
			if less == 0 {
				if float64(weight[u]-weight[x]) >= halfCost {
					ans = x
					high = mid - 1
				} else {
					low = mid + 1
				}
			} else {
				if float64(weight[u]-weight[x]) <= halfCost {
					ans = x
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
		}
		return ans
	}
	var findMedian func(u, v int) int
	findMedian = func(u, v int) int {
		lca := getLca(u, v)
		totalCost := weight[u] + weight[v] - 2*weight[lca]
		if weight[u] == weight[v] {
			return lca
		}
		if weight[u] > weight[v] {
			return findBetween(u, lca, float64(totalCost)/2, 0)
		}
		return findBetween(v, lca, float64(totalCost)/2, 1)
	}
	for i, q := range queries {
		ret[i] = findMedian(q[0], q[1])
	}
	return ret
}
