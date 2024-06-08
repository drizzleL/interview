package main

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	dict := map[int]int{}
	dict[0] = 1
	var ret int
	var presums int
	for _, num := range nums {
		if num%modulo == k {
			presums++
		}
		presums %= modulo
		idx := (presums - k + modulo) % modulo
		ret += dict[idx]
		dict[presums]++
	}
	return int64(ret)
}
