package main

func minOperations18(nums []int) int {
	var ret int
	var flag int
	for _, v := range nums {
		if v^flag == 1 {
			continue
		}
		ret += 1
		flag ^= 1
	}
	return ret
}
