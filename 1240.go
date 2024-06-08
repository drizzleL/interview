package main

// TODO
func tilingRectangle(n int, m int) int {
	cache := map[[2]int]int{}
	var helper func(a, b int) int
	helper = func(a, b int) (ret int) {
		if c, ok := cache[[2]int{a, b}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{a, b}] = ret
		}()
		if a == 0 && b == 0 {
			return 0
		}
		if a == b {
			return 1
		}
		if a > b {
			a, b = b, a
		}
		size1 := 1 + helper(a, b-a)
		size2 := 0
		return min(size1, size2)
	}
	return helper(n, m)
}
