package main

func treeQueries2(n int, edges [][]int, queries [][]int) []int {
	dict := make([][][2]int, n+1)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], [2]int{ed[1], ed[2]})
		dict[ed[1]] = append(dict[ed[1]], [2]int{ed[0], ed[2]})
	}
	to := map[[2]int]int{}
	start, end := make([]int, n+1), make([]int, n+1)
	var dfs func(i, p int, idx int) int
	dfs = func(i, p int, idx int) int {
		start[i] = idx
		for _, child := range dict[i] {
			j := child[0]
			if j == p {
				continue
			}
			to[[2]int{i, j}] = child[1]
			idx += 1
			idx = dfs(j, i, idx)
		}
		end[i] = idx
		return idx
	}
	dfs(1, 0, 0)
	getKey := func(i, j int) ([2]int, int) {
		key := [2]int{i, j}
		old, ok := to[[2]int{i, j}]
		if ok {
			return key, old
		}
		key = [2]int{j, i}
		return key, to[key]
	}
	f := NewFenwick(n)
	var dfs2 func(i, p int)
	dfs2 = func(i, p int) {
		for _, child := range dict[i] {
			j := child[0]
			if j == p {
				continue
			}
			f.add(start[j], child[1])
			f.add(end[j]+1, -child[1])
			dfs2(j, i)
		}
	}
	dfs2(1, 0)
	var ret []int
	for _, q := range queries {
		if q[0] == 1 { // update
			i, j, w := q[1], q[2], q[3]
			key, val := getKey(i, j)
			diff := w - val
			f.add(start[j], diff)
			f.add(end[j]+1, -diff)
			to[key] = w
			continue
		}
		i := q[1]
		ret = append(ret, f.sum(start[i]))
	}
	return ret
}
