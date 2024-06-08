package main

func singleNumber(nums []int) []int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	xor &= -xor
	ret := make([]int, 2)
	for _, num := range nums {
		if num&xor != 0 {
			ret[0] ^= num
		} else {
			ret[1] ^= num
		}
	}
	return ret
}
