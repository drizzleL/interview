package main

func dieSimulator(n int, rollMax []int) int {
	cache := map[[2]int]int{}
	var helper func(k int, last int) int
	helper = func(k int, last int) (ret int) {
		if c, ok := cache[[2]int{k, last}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{k, last}] = ret
		}()
		if k == n {
			return 1
		}
		for i := 0; i < 6; i++ {
			if i == last {
				continue
			}
			for m := 1; m <= rollMax[i] && m+k <= n; m++ {
				ret += helper(k+m, i)
				ret %= 1e9 + 7
			}
		}
		return
	}
	return helper(0, -1)
}
