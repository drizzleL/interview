package main

func findWordsContaining(words []string, x byte) []int {
	var ret []int
	for i, w := range words {
		for _, c := range w {
			if byte(c) == x {
				ret = append(ret, i)
				break
			}
		}
	}
	return ret
}
