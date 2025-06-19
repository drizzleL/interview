package main

func colorTheArray(n int, queries [][]int) []int {
	colors := make([]int, n)
	ret := make([]int, len(queries))
	var cnt int
	check := func(i int, color1, color2 int) int {
		if i < 0 || i >= n {
			return 0
		}
		var ret int
		if colors[i] != 0 && color1 != 0 && colors[i] == color1 {
			ret -= 1
		}
		if colors[i] != 0 && color2 != 0 && colors[i] == color2 {
			ret += 1
		}
		return ret
	}
	for i, q := range queries {
		cnt += check(q[0]-1, colors[q[0]], q[1])
		cnt += check(q[0]+1, colors[q[0]], q[1])
		ret[i] = cnt
		colors[q[0]] = q[1]
	}
	return ret
}
