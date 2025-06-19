package main

func vowelStrings(words []string, queries [][]int) []int {
	sums := make([]int, len(words)+1)
	check := func(c byte) bool {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}
		return false
	}
	for i, w := range words {
		sums[i+1] = sums[i]
		if check(w[0]) && check(w[len(w)-1]) {
			sums[i+1] += 1
		}
	}
	ret := make([]int, len(queries))
	for i, q := range queries {
		a, b := q[0], q[1]
		ret[i] = sums[b+1] - sums[a]
	}
	return ret
}
