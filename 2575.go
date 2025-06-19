package main

func divisibilityArray(word string, m int) []int {
	var left int
	ret := make([]int, len(word))
	for i := range word {
		v := left*10 + int(word[i]-'0')
		left = v % m
		if left != 0 {
			ret[i] = 1
		}
	}
	return ret
}
