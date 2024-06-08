package main

func findCenter(edges [][]int) int {
	dict := map[int]int{}
	for _, ed := range edges {
		dict[ed[0]] += 1
		dict[ed[1]] += 1
	}
	var ret int
	for k, v := range dict {
		if ret == 0 || v > dict[ret] {
			ret = k
		}
	}
	return ret
}
