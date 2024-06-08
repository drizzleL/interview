package main

func distinctDifferenceArray(nums []int) []int {
	suffix := map[int]int{}
	for _, num := range nums {
		suffix[num]++
	}
	prefix := map[int]int{}
	ret := make([]int, len(nums))
	for i, num := range nums {
		prefix[num]++
		suffix[num]--
		if suffix[num] == 0 {
			delete(suffix, num)
		}
		ret[i] = len(prefix) - len(suffix)
	}
	return ret
}
