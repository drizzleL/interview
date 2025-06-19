package main

var cache, upper = 0, 1

func rand10() int {
	for upper < 1e9 {
		cache *= 7
		cache += rand7() - 1
		upper *= 7
	}
	ret := cache % 10
	upper /= 10
	cache /= 10
	return ret
}
func rand7() int {
	return 0
}
