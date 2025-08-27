package main

func minSensors(n int, m int, k int) int {
	return (n + 2*k) / (2*k + 1) * (m + 2*k) / (2*k + 1)
}
