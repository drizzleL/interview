package main

import "strconv"

func crackNumber(ciphertext int) int {
	a, b := 0, 1
	str := strconv.Itoa(ciphertext)
	for i := 1; i < len(str); i++ {
		a, b = b, a+b
		num, _ := strconv.Atoi(str[i-1 : i+1])
		if num > 25 || strconv.Itoa(num) != str[i-1:i+1] {
			a = 0
		}
	}
	return a + b
}
