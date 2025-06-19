package main

func findFinalValue(nums []int, original int) int {
	dict := map[int]bool{}
	for _, num := range nums {
		dict[num] = true
	}
	for dict[original] {
		original *= 2
	}
	return original
}
