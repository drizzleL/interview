package main

func countConsistentStrings(allowed string, words []string) int {
	dict := map[rune]bool{}
	for _, c := range allowed {
		dict[c] = true
	}
	var ret int
	for _, w := range words {
		var flag bool
		for _, c := range w {
			if !dict[c] {
				flag = true
				break
			}
		}
		if !flag {
			ret += 1
		}
	}
	return ret
}
