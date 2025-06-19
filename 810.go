package main

func xorGame(nums []int) bool {
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	return xor == 0 || len(nums)%2 == 0
}
