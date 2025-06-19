package main

func findArray(pref []int) []int {
	ret := make([]int, len(pref))
	ret[0] = pref[0]
	for i := 1; i < len(pref); i++ {
		ret[i] = ret[i-1] ^ pref[i]
	}
	return ret
}
