package main

func evenNumberBitwiseORs(nums []int) int {
	var ret int
	for _, num := range nums {
		if num%2 != 0 {
			continue
		}
		ret |= num
	}
	return ret
}
