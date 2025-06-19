package main

import "log"

func punishmentNumber(n int) int {
	var ret int
	var check func(v, x int) bool
	check = func(v, x int) bool {
		if v == x {
			return true
		}
		base := 10
		for x > 0 && v > x {
			v2 := v % base
			if v2 > x {
				break
			}
			if check(v/base, x-v2) {
				log.Println(v, x, v2)
				return true
			}
			base *= 10
		}
		return false
	}
	for i := 1; i <= n; i++ {
		if check(i*i, i) {
			log.Println(i*i, i)
			ret += i * i
		}
	}
	return ret
}
