package main

func waysToStep(n int) int {
	if n <= 1 {
		return 1
	}
	a, b, c := 1, 1, 2
	for i := 2; i < n; i++ {
		a, b, c = b, c, (a+b+c)%1000000007
	}
	return c
}
