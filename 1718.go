package main

func constructDistancedSequence(n int) []int {
	dict := make([]int, n+1)
	dict[1] = 1
	for i := 2; i <= n; i++ {
		dict[i] += 2
	}
	ret := make([]int, n*2-1)
	var helper func(i int) bool
	helper = func(i int) bool {
		if i == len(ret) {
			return true
		}
		if ret[i] != 0 {
			return helper(i + 1)
		}
		for j := n; j >= 1; j-- {
			if dict[j] == 0 {
				continue
			}
			if j != 1 && (i+j >= len(ret) || ret[i+j] != 0) {
				continue
			}
			ret[i] = j
			dict[j] -= 1
			if j != 1 {
				dict[j] -= 1
				ret[i+j] = j
			}
			if helper(i + 1) {
				return true
			}
			dict[j] += 1
			ret[i] = 0
			if j != 1 {
				dict[j] += 1
				ret[i+j] = 0
			}
		}
		return false
	}
	helper(0)
	return ret
}
