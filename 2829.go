package main

func minimumSum2(n int, k int) int {
	var ret int
	dict := make([]bool, 51)
	var cnt int
	for i := 1; cnt < n; i++ {
		if dict[k-i] {
			continue
		}
		cnt += 1
		dict[i] = true
		ret += i
	}
	return ret
}
