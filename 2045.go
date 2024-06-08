package main

func secondMinimum(n int, edges [][]int, time int, change int) int {
	dict := map[int][]int{}
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	find := func() int {
		nodes := []int{1}
		vis := make([]int, n+1)
		visCnt := make([]int, n+1)
		vis[1] = 0
		visCnt[1] = 1
		for step := 1; len(nodes) != 0; step++ {
			var next []int
			for _, no := range nodes {
				for _, child := range dict[no] {
					if vis[child] == step || visCnt[child] == 2 {
						continue
					}
					vis[child] = step
					visCnt[child] += 1
					if child == n && visCnt[child] == 2 {
						return step
					}
					next = append(next, child)
				}
			}
			nodes = next
		}
		return vis[n]
	}
	step := find()
	var ret int
	changeTerm := change * 2
	for ; step != 0; step-- {
		ret += time
		if step == 1 {
			break
		}
		t := ret / changeTerm
		if ret >= t*changeTerm+change { // can't go
			ret = (t + 1) * changeTerm
		}
	}
	return ret
}
