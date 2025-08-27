package main

func longestSpecialPath(edges [][]int, nums []int) []int {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	valDict := make([]int, maxVal+1)
	for i := range valDict {
		valDict[i] = -1
	}
	dict := make([][][2]int, len(nums))
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], [2]int{ed[1], ed[2]})
		dict[ed[1]] = append(dict[ed[1]], [2]int{ed[0], ed[2]})
	}
	ret := []int{0, 0}
	check := func(start int, presums []int) {
		v := presums[len(presums)-1]
		v -= presums[start]
		if v < ret[0] {
			return
		}
		if v > ret[0] {
			ret[0] = v
			ret[1] = len(presums) - start
			return
		}
		ret[1] = min(ret[1], len(presums)-start)
	}
	checkDup := func(v int, start int) bool {
		if valDict[v] == -1 {
			return false
		}
		lastIdx := valDict[v]
		return lastIdx >= start
	}
	var dfs func(i int, pre int, preDupVal int, preDupIdx int, presums []int, start int)
	dfs = func(i int, pre int, preDupVal, preDupIdx int, presums []int, start int) {
		v := nums[i]
		if checkDup(v, start) {
			if preDupVal == -1 {
				preDupVal = v
				preDupIdx = valDict[v]
			} else {
				currDup := valDict[v]
				if preDupIdx < currDup { // use curr
					start = preDupIdx + 1
					preDupVal = v
					preDupIdx = currDup
				} else {
					start = currDup + 1
				}
			}
		}
		check(start, presums)
		oldIdx := valDict[v]
		valDict[v] = len(presums) - 1
		for _, next := range dict[i] {
			j, l := next[0], next[1]
			if j == pre {
				continue
			}
			dfs(j, i, preDupVal, preDupIdx, append(presums, presums[len(presums)-1]+l), start)
		}
		valDict[v] = oldIdx
	}
	presums := make([]int, 1, len(nums)+1)
	dfs(0, -1, -1, -1, presums, 0)
	return ret
}
