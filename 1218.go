package main

func longestSubsequence(arr []int, difference int) int {
	dict := map[int]int{}
	var ret int
	for _, v := range arr {
		dict[v] = max(dict[v], dict[v-difference]+1)
		ret = max(ret, dict[v])
	}
	return ret
}
