package main

func tribonacci(n int) int {
	v := [3]int{0, 1, 1}
	if n <= 2 {
		return v[n]
	}
	for i := 2; i < n; i++ {
		v[0], v[1], v[2] = v[1], v[2], v[0]+v[1]+v[2]
	}
	return v[2]
}
