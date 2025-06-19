package main

func criticalConnections(n int, connections [][]int) [][]int {
	dict := make([][]int, n)
	for _, conn := range connections {
		dict[conn[0]] = append(dict[conn[0]], conn[1])
		dict[conn[1]] = append(dict[conn[1]], conn[0])
	}
	ranks := make([]int, n)
	for i := range ranks {
		ranks[i] = -1
	}
	cycle := map[[2]int]bool{}
	addCycle := func(i, j int) {
		if i > j {
			i, j = j, i
		}
		cycle[[2]int{i, j}] = true
	}
	checkCycle := func(i, j int) bool {
		if i > j {
			i, j = j, i
		}
		return cycle[[2]int{i, j}]
	}
	var dfs func(i int, rank int) int
	dfs = func(i int, rank int) (ret int) {
		if ranks[i] > 0 {
			return ranks[i]
		}
		ranks[i] = rank
		ret = n
		for _, child := range dict[i] {
			if ranks[child] == rank-1 { // no trace back
				continue
			}
			childRet := dfs(child, rank+1)
			if childRet <= rank {
				addCycle(i, child)
			}
			ret = min(ret, childRet)
		}
		return ret
	}
	dfs(0, 1)
	var ret [][]int
	for _, conn := range connections {
		if checkCycle(conn[0], conn[1]) {
			continue
		}
		ret = append(ret, conn)
	}
	return ret
}
