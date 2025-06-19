package main

func sumOfGoodSubsequences(nums []int) int {
	cntDict := map[int]int{}
	sumDict := map[int]int{}
	var ret int
	for _, num := range nums {
		cntDict[num] += 1 + cntDict[num-1] + cntDict[num+1]
		cntDict[num] %= 1e9 + 7
		sumDict[num] += num + cntDict[num-1]*num + sumDict[num-1] + cntDict[num+1]*num + sumDict[num+1]
		sumDict[num] %= 1e9 + 7
	}
	for _, sum := range sumDict {
		ret += sum
		ret %= 1e9 + 7
	}
	return ret
}
