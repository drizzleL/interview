package main

func differenceOfSums(n int, m int) int {
	sum := (1 + n) * n / 2
	var num2 int
	for i := 1; i*m <= n; i++ {
		num2 += m * i
	}
	return sum - num2 - num2
}
