package main

func maximumLength4(nums []int) int {
	var evenCnt, oddCnt int
	var wiggleCnt int
	flag := nums[0] % 2
	for _, num := range nums {
		if num%2 == 0 {
			evenCnt += 1
		} else {
			oddCnt += 1
		}
		if num%2 == flag {
			wiggleCnt += 1
			flag ^= 1
		}
	}
	return max(wiggleCnt, max(evenCnt, oddCnt))
}
