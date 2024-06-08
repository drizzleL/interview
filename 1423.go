package main

func maxScore(cardPoints []int, k int) int {
	left := make([]int, k+1)
	right := make([]int, k+1)
	for i := 0; i < k; i++ {
		left[i+1] = left[i] + cardPoints[i]
		j := len(cardPoints) - 1 - i
		right[i+1] = right[i] + cardPoints[j]
	}
	var ret int
	for i := 0; i <= k; i++ {
		ret = max(ret, left[i]+right[k-i])
	}
	return ret
}
