package main

func minOperations(nums []int) int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num] += 1
	}
	var ret int
	for _, v := range dict {
		if v == 1 {
			return -1
		}
		ret += (v-1)/3 + 1
	}
	return ret
}
