package main

import "strconv"

func nextGreaterElement(n int) int {
	b := []byte(strconv.Itoa(n))
	var dict [10]int
	for i := len(b) - 1; i >= 0; i-- {
		d := int(b[i] - '0')
		dict[d] += 1
		if i == len(b)-1 || b[i] >= b[i+1] { // found
			continue
		}
		for j := d + 1; j < len(dict); j++ {
			if dict[j] != 0 {
				dict[j] -= 1
				b[i] = byte('0' + j)
				break
			}
		}
		for j, idx := i+1, 0; j < len(b); j++ {
			for dict[idx] == 0 {
				idx += 1
			}
			b[j] = byte('0' + idx)
			dict[idx] -= 1
		}
		ret, _ := strconv.Atoi(string(b))
		return ret
	}
	return -1
}
