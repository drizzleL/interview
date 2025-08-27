package main

import "strconv"

func maximum69Number(num int) int {
	b := []byte(strconv.Itoa(num))
	for i := 0; i < len(b); i++ {
		if b[i] == '6' {
			b[i] = '9'
			break
		}
	}
	ret, _ := strconv.Atoi(string(b))
	return ret
}
