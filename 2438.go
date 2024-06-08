package main

func productQueries(n int, queries [][]int) []int {
	powers := []int{}
	base := 1
	for n != 0 {
		if n&1 != 0 {
			powers = append(powers, base)
		}
		n >>= 1
		base <<= 1
	}
	ret := make([]int, len(queries))
	for i, q := range queries {
		v := 1
		for j := q[0]; j <= q[1]; j++ {
			v *= powers[j]
			v %= 1e9 + 7
		}
		ret[i] = v
	}
	return ret
}
