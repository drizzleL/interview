package main

func findClosest(x int, y int, z int) int {
	a, b := abs(x-z), abs(y-z)
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	return 2
}
