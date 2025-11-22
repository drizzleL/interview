package main

func minimumOperations13(nums []int) int {
	var ret int
	for _, num := range nums {
		if num%3 == 0 {
			continue
		}
		ret += 1
	}
	return ret
}
