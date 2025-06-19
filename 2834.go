package main

func minimumPossibleSum(n int, target int) int {
	pre := target / 2
	if pre > n {
		return ((1 + n) * n / 2) % (1e9 + 7)
	}
	preSum := (1 + pre) * pre / 2
	afterSize := n - pre
	afterSum := (target + target + afterSize - 1) * afterSize / 2
	return (preSum + afterSum) % (1e9 + 7)
}
