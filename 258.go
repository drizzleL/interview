package main

func addDigits(num int) int {
	for num >= 10 {
		var newnum int
		for num != 0 {
			newnum += num % 10
			num /= 10
		}
		num = newnum
	}
	return num
}
