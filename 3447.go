package main

func assignElements(groups []int, elements []int) []int {
	var maxVal int
	for _, g := range groups {
		maxVal = max(maxVal, g)
	}
	idxs := make([]int, maxVal+1)
	for i := range idxs {
		idxs[i] = -1
	}
	for i, ele := range elements {
		if ele >= len(idxs) || idxs[ele] != -1 {
			continue
		}
		for j := 1; j*ele < len(idxs); j++ {
			if idxs[j*ele] != -1 {
				continue
			}
			idxs[j*ele] = i
		}
	}
	ret := make([]int, 0, len(groups))
	for _, g := range groups {
		ret = append(ret, idxs[g])
	}
	return ret
}
