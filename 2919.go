package main

func minIncrementOperations(nums []int, k int) int64 {
	var a, b, c int
	for _, num := range nums {
		a, b, c = b, c, min(min(a, b), c)+max(0, k-num)
	}
	return int64(min(min(a, b), c))
}
