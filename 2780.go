package main

func minimumIndex(nums []int) int {
	dict := map[int]int{}
	var maxVal int
	for _, num := range nums {
		dict[num] += 1
		if dict[num] > dict[maxVal] {
			maxVal = num
		}
	}
	var preCnt int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == maxVal {
			preCnt += 1
		}
		if preCnt*2 > i+1 && (dict[maxVal]-preCnt)*2 > len(nums)-i-1 {
			return i
		}
	}
	return -1
}
