package main

func numberOf2sInRange(n int) int {
	var total int
	base := 1
	var k int
	var ret int
	for n != 0 {
		d := n % 10
		if d != 0 {
			tmp := base
			if d == 2 {
				tmp = k + 1
			} else if d == 1 {
				tmp = 0
			}
			ret += total*d + tmp
		}
		k += base * d
		total += total*9 + base
		base *= 10
		n /= 10
	}
	return ret
}
