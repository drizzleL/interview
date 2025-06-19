package main

func largestCombination(candidates []int) int {
	dict := [32]int{}
	addSum := func(v int) {
		for i := 0; i < 32; i++ {
			dict[i] += v & 1
			v >>= 1
		}
	}
	var ret int
	for i := 0; i < len(candidates); i++ {
		addSum(candidates[i])
	}
	for _, v := range dict {
		ret = max(ret, v)
	}
	return ret
}
