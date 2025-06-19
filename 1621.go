package main

func numberOfSets3(n int, k int) int {
	return combination(n+k-1, 2*k)
}
func numberOfSets(n int, k int) int {
	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, k+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var helper func(i, j int) int
	helper = func(i, j int) (ret int) {
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		defer func() {
			cache[i][j] = ret
		}()
		if i-1 == j {
			return 1
		}
		if j == 1 {
			return (i - 1) * i / 2
		}
		ret += helper(i-1, j)
		for m := 1; m <= i-j; m++ {
			ret += helper(i-m, j-1)
		}
		ret %= 1e9 + 7
		return
	}
	return helper(n, k)
}
