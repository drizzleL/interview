package main

func numWays2(words []string, target string) int {
	dict := make([][26]int, len(words[0]))
	for _, w := range words {
		for i, c := range w {
			dict[i][c-'a'] += 1
		}
	}
	cache := make([][]int, len(dict))
	for i := range cache {
		cache[i] = make([]int, len(target))
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var helper func(i, j int) int
	helper = func(i, j int) (ret int) {
		if j == len(target) {
			return 1
		}
		if i == len(dict) {
			return 0
		}
		if c := cache[i][j]; c != -1 {
			return c
		}
		defer func() {
			cache[i][j] = ret
		}()
		c := target[j] - 'a'
		if dict[i][c] != 0 {
			ret += dict[i][c] * helper(i+1, j+1)
		}
		ret += helper(i+1, j)
		ret %= 1e9 + 7
		return
	}
	return helper(0, 0)
}
