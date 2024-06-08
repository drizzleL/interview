package main

func minSteps(s string, t string) int {
	dict := [26]int{}
	for _, c := range s {
		dict[c-'a'] += 1
	}
	for _, c := range t {
		dict[c-'a'] -= 1
	}
	var sum int
	for _, v := range dict {
		if v > 0 {
			sum += v
		} else {
			sum -= v
		}
	}
	return (sum + 1) / 2
}
