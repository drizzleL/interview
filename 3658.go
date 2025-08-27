package main

func gcdOfOddEvenSums(n int) int {
	evenSum := (1 + n) * n
	oddSum := n * n
	return gcd(evenSum, oddSum)
}
