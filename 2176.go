package main

func countPairs8(nums []int, k int) int {
	dict := map[int][]int{}
	var ret int
	for i, num := range nums {
		for _, preIdx := range dict[num] {
			if i*preIdx%k == 0 {
				ret += 1
			}
		}
		dict[num] = append(dict[num], i)
	}
	return ret
}
