package main

import "math"

func minMoves88(nums []int, k int) int {
	var r []int
	for i, num := range nums {
		if num == 0 {
			continue
		}
		r = append(r, i-len(r))
	}
	s := make([]int, len(r)+1) // S[0] = 0
	for i := 0; i < len(r); i++ {
		s[i+1] = s[i] + r[i]
	}
	getSum := func(i, j int) int {
		return s[j+1] - s[i]
	}
	ret := math.MaxInt32
	for i := 0; i <= len(r)-k; i++ {
		mid := i + k/2
		presum, suffsum := getSum(i, mid-1), getSum(mid+1, i+k-1)
		precnt, suffcnt := mid-i, i+k-1-mid
		ret = min(ret, (precnt*r[mid]-presum)+(suffsum-suffcnt*r[mid]))
	}
	return ret
}
