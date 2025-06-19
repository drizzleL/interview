package main

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	makeDict := func(edges [][]int) map[int][]int {
		dict := map[int][]int{}
		for _, ed := range edges {
			dict[ed[0]] = append(dict[ed[0]], ed[1])
			dict[ed[1]] = append(dict[ed[1]], ed[0])
		}
		return dict
	}
	dict1 := makeDict(edges1)
	dict2 := makeDict(edges2)
	helper := func(x int, dict map[int][]int, k int) int {
		if k < 0 {
			return 0
		}
		seen := make([]bool, len(dict))
		seen[x] = true
		ret := 1
		nodes := []int{x}
		for i := 0; i < k && len(nodes) != 0; i++ {
			var next []int
			for _, n := range nodes {
				for _, child := range dict[n] {
					if seen[child] {
						continue
					}
					ret += 1
					seen[child] = true
					next = append(next, child)
				}
			}
			nodes = next
		}
		return ret
	}
	var maxEdge2 int
	for i := range dict2 {
		maxEdge2 = max(maxEdge2, helper(i, dict2, k-1))
	}
	ret := make([]int, len(dict1))
	for i := range dict1 {
		ret[i] = helper(i, dict1, k) + maxEdge2
	}
	return ret
}
