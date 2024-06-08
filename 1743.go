package main

func restoreArray(adjacentPairs [][]int) []int {
	dict := map[int][]int{}
	for i := 0; i < len(adjacentPairs); i++ {
		pair := adjacentPairs[i]
		dict[pair[0]] = append(dict[pair[0]], pair[1])
		dict[pair[1]] = append(dict[pair[1]], pair[0])
	}
	var start int
	for k, v := range dict {
		if len(v) == 1 {
			start = k
		}
	}
	seen := map[int]bool{}
	var ret []int
	ret = append(ret, start)
	seen[start] = true
	for len(ret) != len(dict) {
		last := ret[len(ret)-1]
		for _, v := range dict[last] {
			if seen[v] {
				continue
			}
			ret = append(ret, v)
			seen[v] = true
		}
	}
	return ret
}
