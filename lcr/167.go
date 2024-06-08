package main

func dismantlingAction2(arr string) int {
	var ret int
	var start int
	dict := make(map[rune]int)
	for i, c := range arr {
		oldIdx, exist := dict[c]
		if exist && oldIdx >= start {
			start = oldIdx + 1
		}
		dict[c] = i
		if i-start+1 > ret {
			ret = i - start + 1
		}
	}
	return ret
}
