package main

func maximumDetonation(bombs [][]int) int {
	dict := make([][]int, len(bombs))
	for i := 0; i < len(bombs); i++ {
		for j := 0; j < len(bombs); j++ {
			if j == i {
				continue
			}
			a, b := bombs[i][0]-bombs[j][0], bombs[i][1]-bombs[j][1]
			if bombs[i][2]*bombs[i][2] >= a*a+b*b {
				dict[i] = append(dict[i], j)
			}
		}
	}
	check := func(i int) int {
		seen := make([]bool, len(bombs))
		nodes := []int{i}
		var ret int
		for len(nodes) > 0 {
			var next []int
			for _, node := range nodes {
				seen[node] = true
				ret += 1
				for _, child := range dict[node] {
					if seen[child] {
						continue
					}
					seen[child] = true
					next = append(next, child)
				}
			}
			nodes = next
		}
		return ret
	}
	var ret int
	for i := 0; i < len(bombs); i++ {
		ret = max(ret, check(i))
	}
	return ret
}
