package main

import (
	"strconv"
)

func maxDiff(num int) int {
	str := strconv.Itoa(num)
	replaceWith := func(str string, c1, c2 byte) int {
		b := []byte(str)
		for i := 0; i < len(b); i++ {
			var c byte
			if b[i] == c1 || (b[i] == c2 && b[i] != c1) {
				continue
			}
			if i == 0 {
				c = c1
			} else {
				c = c2
			}
			oldC := b[i]
			for j := i; j < len(b); j++ {
				if b[j] == oldC {
					b[j] = c
				}
			}
			break
		}
		num, _ := strconv.Atoi(string(b))
		return num
	}
	return replaceWith(str, '9', '9') - replaceWith(str, '1', '0')
}
