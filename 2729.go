package main

func isFascinating(n int) bool {
	dict := make([]bool, 10)
	fill := func(num int) bool {
		if num > 999 {
			return false
		}
		for num != 0 {
			remainder := num % 10
			if remainder == 0 {
				return false
			}
			if dict[remainder] {
				return false
			}
			dict[remainder] = true
			num /= 10
		}
		return true
	}
	return fill(n) && fill(n*2) && fill(n*3)
}
