package main

func minElements(nums []int, limit int, goal int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return (abs(sum-goal) + limit - 1) / limit
}
