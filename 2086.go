package main

import "strings"

func minimumBuckets(hamsters string) int {
	if len(hamsters) == 0 {
		return 0
	}
	if len(hamsters) == 1 {
		if hamsters[0] == 'H' {
			return -1
		}
		return 0
	}
	if hamsters[:2] == "HH" || hamsters[len(hamsters)-2:] == "HH" {
		return -1
	}
	var hcount, consCount int
	for _, h := range hamsters {
		switch h {
		case 'H':
			hcount += 1
			consCount += 1
			if consCount == 3 {
				return -1
			}
		case '.':
			consCount = 0
		}
	}
	cnt2 := strings.Count(hamsters, "H.H")
	return hcount - cnt2
}
