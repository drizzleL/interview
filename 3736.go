package main

func minMoves9(nums []int) int {
	var ret, maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	for _, num := range nums {
		ret += maxVal - num
	}
	return ret
}
