package main

func isPowerOfFour(n int) bool {
	num := 1
	for num < n {
		num *= 4
	}
	return num == n
}
