package main

func maxProduct(words []string) int {
	dict := make([]int, len(words))
	for i, word := range words {
		for _, c := range word {
			dict[i] |= 1 << (c - 'a')
		}
	}
	var ret int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if dict[i]&dict[j] == 0 {
				tmp := len(words[i]) * len(words[j])
				if tmp > ret {
					ret = tmp
				}

			}
		}
	}
	return ret
}
