package main

func getGoodIndices(variables [][]int, target int) []int {
	powMod := func(a, b int, mod int) int {
		ret := 1
		for b > 0 {
			if b&1 > 0 {
				ret *= a
				ret %= mod
			}
			a *= a
			a %= mod
			b >>= 1
		}
		return ret
	}
	var ret []int
	for i, v := range variables {
		v1 := powMod(v[0], v[1], 10)
		v2 := powMod(v1, v[2], v[3])
		if v2 == target {
			ret = append(ret, i)
		}
	}
	return ret
}
