package main

func longestPalindrome4(words []string) int {
	dict := map[string]int{}
	rev := func(x string) string {
		b := []byte(x)
		b[0], b[1] = b[1], b[0]
		return string(b)
	}
	var ret int
	for _, w := range words {
		revW := rev(w)
		if dict[revW] > 0 {
			ret += 4
			dict[revW] -= 1
			continue
		}
		dict[w] += 1
	}
	for k, v := range dict {
		if v == 0 {
			continue
		}
		if k[0] != k[1] {
			continue
		}
		ret += 2
	}
	return ret
}
