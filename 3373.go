package main

func maxTargetNodes2(edges1 [][]int, edges2 [][]int) []int {
	toDict := func(edges [][]int) [][]int {
		dict := make([][]int, len(edges)+1)
		for _, ed := range edges {
			dict[ed[0]] = append(dict[ed[0]], ed[1])
			dict[ed[1]] = append(dict[ed[1]], ed[0])
		}
		return dict
	}
	var helper func(i int, p int, flag int, dict [][]int, ret []int)
	helper = func(i int, p int, flag int, dict [][]int, ret []int) {
		ret[i] = flag
		for _, child := range dict[i] {
			if child == p {
				continue
			}
			helper(child, i, 1-flag, dict, ret)
		}
	}
	getColorDict := func(edges [][]int) ([]int, []int) {
		colors := make([]int, len(edges)+1)
		helper(0, -1, 0, toDict(edges), colors)
		colorDict := make([]int, 2)
		for _, c := range colors {
			colorDict[c]++
		}
		return colors, colorDict
	}
	c1, dict1 := getColorDict(edges1)
	_, dict2 := getColorDict(edges2)
	cnt2 := max(dict2[0], dict2[1])
	ret := make([]int, len(edges1)+1)
	for i := range ret {
		ret[i] = dict1[c1[i]] + cnt2
	}
	return ret
}
