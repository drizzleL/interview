package main

func sumOfSquares(nums []int) int {
	var ret int
	for i, num := range nums {
		if len(nums)%(i+1) == 0 {
			ret += num * num
		}
	}
	return ret
}
