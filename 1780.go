package main

func checkPowersOfThree(n int) bool {
	for n != 1 {
		if n%3 > 1 {
			return false
		}
		n /= 3
	}
	return true
}
