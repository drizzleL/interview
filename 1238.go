package main

func circularPermutation(n int, start int) []int {
	size := 1 << n
	dict := make([]bool, size)
	ret := make([]int, 0, size)
	for len(ret) != size {
		ret = append(ret, start)
		dict[start] = true
		for i := 0; i < n; i++ {
			flag := 1 << i
			if dict[start^flag] {
				continue
			}
			start ^= flag
			break
		}
	}
	return ret
}
