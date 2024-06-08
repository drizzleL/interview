package main

func countSubstrings(s string, c byte) int64 {
	var sum int
	for _, v := range s {
		if byte(v) == c {
			sum += 1
		}
	}
	return int64((1 + sum) * sum / 2)
}
