package main

func minIncrease(n int, edges [][]int, cost []int) int {
	dict := make([][]int, n)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	var ret int
	var helper func(x int, p int) int
	helper = func(x int, p int) int {
		var maxCost int
		var cnt int
		var maxCostCnt int
		for _, next := range dict[x] {
			if next == p {
				continue
			}
			cnt += 1
			childCost := helper(next, x)
			if childCost > maxCost {
				maxCost = childCost
				maxCostCnt = 1
			} else if childCost == maxCost {
				maxCostCnt += 1
			}
		}
		ret += cnt - maxCostCnt
		return maxCost + cost[x]
	}
	helper(0, -1)
	return ret
}
