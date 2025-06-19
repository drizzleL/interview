package main

func findMaxLength(nums []int) int {
	dict := map[int]int{}
	dict[0] = -1
	var cnt [2]int
	var ret int
	for i, num := range nums {
		cnt[num] += 1
		diff := cnt[0] - cnt[1]
		old, ok := dict[diff]
		if !ok {
			dict[diff] = i
			continue
		}
		ret = max(ret, i-old)
	}
	return ret
}
