package main

func findCircleNum(isConnected [][]int) int {
	var ret int
	vis := make([]bool, len(isConnected))
	var infect func(i int)
	infect = func(i int) {
		for j := 0; j < len(isConnected); j++ {
			if isConnected[i][j] == 0 {
				continue
			}
			if vis[j] || i == j {
				continue
			}
			vis[j] = true
			infect(j)
		}
	}
	for i := 0; i < len(isConnected); i++ {
		if vis[i] {
			continue
		}
		ret++
		infect(i)
	}
	return ret
}

func findCircleNum2(isConnected [][]int) int {
	parent := make([]int, len(isConnected))
	for i := range parent {
		parent[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return i
	}
	union := func(i, j int) {
		parent[find(i)] = find(j)
	}
	for i := 0; i < len(parent); i++ {
		for j := i + 1; j < len(parent); j++ {
			if isConnected[i][j] == 0 {
				continue
			}
			union(i, j)
		}
	}
	var ret int
	for i, v := range parent {
		if i == v {
			ret++
		}
	}
	return ret
}
