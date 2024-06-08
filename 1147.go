package main

func longestDecomposition(text string) int {
	l, r := 0, len(text)-1
	var ret int
	for l <= r {
		lend, rstart := l, r
		var flag bool
		for ; lend < rstart; lend, rstart = lend+1, rstart-1 {
			if text[l:lend+1] == text[rstart:r+1] {
				ret += 2
				flag = true
				break
			}
		}
		if !flag {
			ret += 1
			break
		}
		l, r = lend+1, rstart-1
	}
	return ret
}
