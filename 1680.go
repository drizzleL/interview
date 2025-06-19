package main

func concatenatedBinary(n int) int {
	var ret int
	for i := 1; i <= n; i++ {
		for j := i; j > 0; j >>= 1 {
			ret <<= 1
			ret %= 1e9 + 7
		}
		ret += i
		ret %= 1e9 + 7
	}
	return ret
}
