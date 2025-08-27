package main

func checkDivisibility(n int) bool {
	var sum int
	last := 1
	for n2 := n; n2 != 0; n2 /= 10 {
		d := n2 % 10
		sum += d
		last *= d
	}
	sum += last
	return n%sum == 0
}
