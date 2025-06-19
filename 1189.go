package main

import "math"

func maxNumberOfBalloons(text string) int {
	var dict [26]int
	for _, c := range text {
		dict[c-'a'] += 1
	}
	ret := math.MaxInt32
	for _, c := range "ban" {
		ret = min(ret, dict[c-'a'])
	}
	for _, c := range "lo" {
		ret = min(ret, dict[c-'a']/2)
	}
	if ret == math.MaxInt32 {
		return 0
	}
	return ret
}
