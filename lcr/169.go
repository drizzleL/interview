package main

func dismantlingAction(arr string) byte {
	dict := map[rune]int{}
	for _, c := range arr {
		dict[c]++
	}
	for _, c := range arr {
		if dict[c] == 1 {
			return byte(c)
		}
	}
	return 0
}
