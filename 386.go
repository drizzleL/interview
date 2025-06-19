package main

func lexicalOrder(n int) []int {
	var ret []int
	var helper func(pre int)
	helper = func(pre int) {
		if pre > n {
			return
		}
		if pre != 0 {
			ret = append(ret, pre)
		}
		if pre != 0 {
			helper(pre * 10)
		}
		for i := 1; i <= 9; i++ {
			helper(pre*10 + i)
		}
	}
	helper(0)
	return ret
}
