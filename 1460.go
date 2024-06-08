package main

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	dict := map[int]int{}
	for _, num := range arr {
		dict[num] += 1
	}
	for _, num := range target {
		dict[num] -= 1
		if dict[num] < 0 {
			return false
		}
	}
	return true
}
