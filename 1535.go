package main

func getWinner(arr []int, k int) int {
	var winnerIdx int
	var winnerCnt int
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[winnerIdx] {
			winnerIdx = i
			winnerCnt = 1
		} else {
			winnerCnt += 1
		}
		if winnerCnt == k {
			break
		}
	}
	return arr[winnerIdx]
}
