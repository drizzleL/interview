package main

import "sort"

func maxConsecutive(bottom int, top int, special []int) int {
	var ret int
	sort.Ints(special)
	ret = max(ret, special[0]-bottom)
	ret = max(ret, top-special[len(special)-1])
	for i := 1; i < len(special); i++ {
		ret = max(ret, special[i]-special[i-1])
	}
	return ret
}
