package main

func wordCount(startWords []string, targetWords []string) int {
	dict := map[int]bool{}
	wordToInt := func(word string) int {
		var dest int
		for _, c := range word {
			bit := 1 << int(c-'a')
			dest |= bit
		}
		return dest
	}
	for _, t := range targetWords {
		dict[wordToInt(t)] = false
	}
	for _, s := range startWords {
		dest := wordToInt(s)
		for i := 0; i < 26; i++ {
			bit := 1 << i
			if dest&bit != 0 {
				continue
			}
			if _, ok := dict[dest|bit]; ok {
				dict[dest|bit] = true
			}
		}
	}
	var ret int
	for _, t := range targetWords {
		if dict[wordToInt(t)] {
			ret += 1
		}
	}
	return ret
}
