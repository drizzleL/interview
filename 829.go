package main

func consecutiveNumbersSum(n int) int {
	var ret int
	for i := 1; ; i++ {
		num := (n * 2)
		if num%i != 0 {
			continue
		}
		num = num/i - i + 1
		if num <= 0 {
			break
		}
		if num%2 == 0 {
			ret += 1
		}
	}
	return ret
}
