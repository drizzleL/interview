package main

func findMaxK(nums []int) int {
	var k int
	dict := map[int]bool{}
	for _, num := range nums {
		dict[num] = true
		if dict[-num] {
			k = max(k, abs(num))
		}
	}
	if k == 0 {
		return -1
	}
	return k
}
