package main

func minimizeConcatenatedLength(words []string) int {
	cache := map[[3]int]int{}
	var helper func(i int, head, tail int) int
	helper = func(i int, head, tail int) (ret int) {
		if i == len(words) {
			return 0
		}
		if c, ok := cache[[3]int{i, head, tail}]; ok {
			return c
		}
		defer func() {
			cache[[3]int{i, head, tail}] = ret
		}()
		w := words[i]
		a := len(w) + helper(i+1, head, int(w[len(w)-1]-'a'))
		if tail == int(w[0]-'a') {
			a -= 1
		}
		b := len(w) + helper(i+1, int(w[0]-'a'), tail)
		if head == int(w[len(w)-1]-'a') {
			b -= 1
		}
		return min(a, b)
	}
	return len(words[0]) + helper(1, int(words[0][0]-'a'), int(words[0][len(words[0])-1]-'a'))
}
