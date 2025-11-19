package main

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	ret := make([]int, len(parents))
	for i := range ret {
		ret[i] = 1
	}
	children := make([][]int, len(parents))
	for i, p := range parents {
		if p == -1 {
			continue
		}
		children[p] = append(children[p], i)
	}
	var node, maxVal int
	for i, num := range nums {
		if num == 1 {
			node = i
		}
		maxVal = max(maxVal, num)
	}
	seen := make([]bool, maxVal+2)
	idx := 1
	getMissing := func() int {
		for seen[idx] {
			idx += 1
		}
		return idx
	}
	var mark func(i int)
	mark = func(i int) {
		seen[nums[i]] = true
		for _, c := range children[i] {
			mark(c)
		}
	}
	for _, c := range children[node] {
		mark(c)
	}
	for node != -1 {
		seen[nums[node]] = true
		ret[node] = getMissing()
		if parents[node] == -1 {
			break
		}
		for _, c := range children[parents[node]] {
			if c == node {
				continue
			}
			mark(c)
		}
		node = parents[node]
	}
	return ret
}
