package main

func totalMoney(n int) int {
	var ret int
	weeks := n / 7
	weekMoney := (1 + 7) * 7 / 2
	ret += weekMoney*weeks + weeks*(weeks-1)/2*7
	days := n % 7
	for i := 0; i < days; i++ {
		ret += weeks + i + 1
	}
	return ret
}
