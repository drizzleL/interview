package main

func minCostToEqualizeArray(nums []int, cost1 int, cost2 int) int {
	if len(nums) == 0 {
		return 0
	}
	var sum int
	minNum, maxNum := nums[0], nums[0]
	for _, num := range nums {
		sum += num
		maxNum = max(maxNum, num)
		minNum = min(minNum, num)
	}
	if cost1*2 <= cost2 || len(nums) <= 2 {
		return maxNum*len(nums) - sum
	}
	return 0

}
