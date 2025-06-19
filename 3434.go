package main

func maxFrequency4(nums []int, k int) int {
	dict := map[int]int{}
	var ret int
	for _, num := range nums {
		dict[num] = max(dict[num], dict[k]) + 1
		ret = max(ret, dict[num]-dict[k])
	}
	return ret + dict[k]
}
