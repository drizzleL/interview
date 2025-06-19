package main

func maxProduct6(n int) int {
	var a, b int
	for n != 0 {
		c := n % 10
		n /= 10
		if c >= a {
			a, b = c, a
			continue
		}
		if c >= b {
			b = c
		}
	}
	return a * b
}
