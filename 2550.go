package main

func monkeyMove(n int) int {
	v := fastPow(2, n, 1)
	v -= 2
	if v <= 0 {
		v += 1e9 + 7
	}
	return v
}
