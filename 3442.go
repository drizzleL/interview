package main

func maxDifference(s string) int {
	var dict [26]int
	for _, c := range s {
		dict[c-'a']++
	}
	oddMax, evenMin := 0, len(s)
	for _, v := range dict {
		if v == 0 {
			continue
		}
		if v%2 == 0 {
			evenMin = min(evenMin, v)
		} else {
			oddMax = max(oddMax, v)
		}
	}
	return oddMax - evenMin
}
