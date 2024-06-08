package main

func findLonely(nums []int) []int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num] += 1
	}
	var ret []int
	for k, v := range dict {
		if v > 1 {
			continue
		}
		if dict[k-1] != 0 || dict[k+1] != 0 {
			continue
		}
		ret = append(ret, k)
	}
	return ret
}
