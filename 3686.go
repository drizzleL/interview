package main

func countStableSubsequences(nums []int) int {
	var even, odd [2]int
	mod := 1_000_000_007
	var ret int
	for _, num := range nums {
		switch num % 2 {
		case 0: // even
			ret += even[0] + odd[0] + odd[1] + 1
			even[1] += even[0]
			even[1] %= mod
			even[0] += odd[0] + odd[1] + 1
			even[0] %= mod
		case 1: // odd
			ret += odd[0] + even[0] + even[1] + 1
			odd[1] += odd[0]
			odd[1] %= mod
			odd[0] += even[0] + even[1] + 1
			odd[0] %= mod
		}
		ret %= mod
	}
	return ret
}
