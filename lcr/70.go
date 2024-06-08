package main

func singleNonDuplicate(nums []int) int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	return xor
}
