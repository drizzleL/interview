package main

import "strconv"

func maximumSwap(num int) int {
	b := []byte(strconv.Itoa(num))
	var idx int
	for i := 1; i < len(b); i++ {
		if b[i-1] >= b[i] {
			continue
		}
		idx = i
		break
	}
	if idx == 0 {
		return num
	}
	for i := idx; i < len(b); i++ {
		if b[i] >= b[idx] {
			idx = i
		}
	}
	for i := 0; i < len(b); i++ {
		if b[i] >= b[idx] {
			continue
		}
		b[i], b[idx] = b[idx], b[i]
		break
	}
	ret, _ := strconv.Atoi(string(b))
	return ret
}
