package main

func minSteps3(s string, t string) int {
	var dict1, dict2 [26]int
	for _, c := range s {
		dict1[c-'a'] += 1
	}
	for _, c := range t {
		dict1[c-'a'] += 1
	}
	var ret int
	for i := range dict1 {
		ret += abs(dict1[i] - dict2[i])
	}
	return ret
}
