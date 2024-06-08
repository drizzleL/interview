package main

func sumZero(n int) []int {
	var ret []int
	if n%2 == 1 {
		ret = append(ret, 0)
		n -= 1
	}
	x := 1
	for n != 0 {
		ret = append(ret, x, -x)
		x += 1
		n -= 2
	}
	return ret
}
