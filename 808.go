package main

func soupServings(n int) float64 {
	cnt := (n + 24) / 25
	cache := map[[2]int]float64{}
	var helper func(a, b int) float64
	helper = func(a, b int) (ret float64) {
		if a <= 0 && b <= 0 {
			return 0.5
		}
		if a <= 0 {
			return 1
		}
		if b <= 0 {
			return 0
		}
		if c, ok := cache[[2]int{a, b}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{a, b}] = ret
		}()
		return 0.25 * (helper(a-4, b) + helper(a-2, b-2) + helper(a-3, b-1) + helper(a-1, b-3))
	}
	var ret float64
	for i := 1; i < cnt; i++ {
		ret = helper(i, i)
		if 1-ret < 1e-5 {
			return 1
		}
	}
	return ret
}
