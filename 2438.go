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

func productQueries2(n int, queries [][]int) []int {
	var powers []int
	for base := 0; n != 0; base++ {
		if n&(1<<base) != 0 {
			powers = append(powers, base)
			n ^= 1 << base
		}
	}
	presums := make([]int, len(powers)+1)
	for i := 1; i < len(presums); i++ {
		presums[i] = presums[i-1] + powers[i-1]
	}
	ret := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		tmp := presums[r+1] - presums[l]
		ret[i] = fastPow(2, tmp, 1)
	}
	return ret
}
