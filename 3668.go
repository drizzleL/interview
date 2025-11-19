package main

func recoverOrder(order []int, friends []int) []int {
	dict := map[int]bool{}
	for _, f := range friends {
		dict[f] = true
	}
	var ret []int
	for _, v := range order {
		if dict[v] {
			ret = append(ret, v)
		}
	}
	return ret
}
