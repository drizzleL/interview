package main

func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	for i := 1; i <= 9; i++ {
		v := i * k
		if v%10 == num%10 && v <= num {
			return i
		}
	}
	return -1
}
