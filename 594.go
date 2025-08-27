package main

func findLHS(nums []int) int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num] += 1
	}
	var ret int
	for k, v := range dict {
		v2, ok := dict[k+1]
		if !ok {
			continue
		}
		ret = max(ret, v+v2)
	}
	return ret
}
