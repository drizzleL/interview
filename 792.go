package main

func numMatchingSubseq(s string, words []string) int {
	var ret int
	check := func(w string) bool {
		var j int
		for i := 0; j < len(w) && i < len(s); i++ {
			if s[i] == w[j] {
				j++
			}
		}
		return j == len(w)
	}
	for _, w := range words {
		if check(w) {
			ret += 1
		}

	}
	return ret
}
