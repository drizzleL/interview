package main

func oddString(words []string) string {
	toDiff := func(word string) [20]int {
		ret := [20]int{}
		for i := 1; i < len(word); i++ {
			ret[i-1] = int(word[i]) - int(word[i-1])
		}
		return ret
	}
	d1 := toDiff(words[0])
	d2 := toDiff(words[1])
	d3 := toDiff(words[2])
	if d1 == d2 && d1 == d3 {
		for i := 3; i < len(words); i++ {
			d := toDiff(words[i])
			if d != d1 {
				return words[i]
			}
		}
	}
	if d1 == d2 {
		return words[2]
	}
	if d1 == d3 {
		return words[1]
	}
	return words[0]
}
