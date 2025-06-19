package main

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	var getDiameter func(edges [][]int) int
	getDiameter = func(edges [][]int) int {
		dict := map[int][]int{}
		for _, ed := range edges {
			dict[ed[0]] = append(dict[ed[0]], ed[1])
			dict[ed[1]] = append(dict[ed[1]], ed[0])
		}
		helper := func(x int) (int, int) {
			seen := make([]bool, len(edges)+1)
			nodes := []int{x}
			seen[x] = true
			for step := 0; len(nodes) != 0; step++ {
				var next []int
				for _, n := range nodes {
					for _, child := range dict[n] {
						if seen[child] {
							continue
						}
						seen[child] = true
						next = append(next, child)
					}
				}
				if len(next) == 0 {
					return nodes[0], step
				}
				nodes = next
			}
			return 0, 0
		}
		k, _ := helper(0)
		_, diameter := helper(k)
		return diameter
	}
	d1, d2 := getDiameter(edges1), getDiameter(edges2)
	ret := max(d1, d2)
	ret = max(ret, (d1+1)/2+(d2+1)/2+1)
	return ret
}
