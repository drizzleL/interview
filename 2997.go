package main

func minOperationsV2(nums []int, k int) int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	var ret int
	for xor != 0 || k != 0 {
		if (xor & 1) != (k & 1) {
			ret += 1
		}
		xor >>= 1
		k >>= 1
	}
	return ret
}
