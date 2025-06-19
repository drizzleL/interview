package main

import (
	"sort"
)

// TODO
func maxPartitionsAfterOperations(s string, k int) int {
	if k == 26 {
		return 1
	}
	var cnt int
	var flag int
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		c := s[i] - 'a'
		if cnt == k && flag&(1<<c) == 0 {
			cnt = 0
			flag = 0
			dp[i] += 1
		}
		if flag&(1<<c) == 0 {
			cnt += 1
		}
		flag |= 1 << c
	}
	idxDict := [26][]int{}
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		idxDict[c] = append(idxDict[c], i)
	}
	flag = 0
	cnt = 0
	var ret int
	var preCnt int
	var findEnd func(i int, cnt int, flag int) int
	findEnd = func(i int, cnt int, flag int) int {
		var extra int
		if cnt == k {
			extra = 1
			cnt = 0
		}
		var arr []int
		candidates := flag
		for i := range idxDict {
			if len(idxDict[i]) == 0 {
				continue
			}
			if cnt != 0 && flag&(1<<i) != 0 {
				continue
			}
			candidates |= 1 << i
			arr = append(arr, idxDict[i][0])
		}
		if candidates == 1<<26-1 {
			return 1 + extra + dp[arr[k-cnt]]
		}
		sort.Ints(arr)
		if len(arr) <= k-cnt-1 {
			return 1 + extra
		}
		return dp[arr[k-cnt-1]] + 1 + extra
	}
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		idxDict[c] = idxDict[c][1:]
		if cnt == k && flag&(1<<c) == 0 { // reset
			cnt = 0
			flag = 0
			preCnt += 1
		}
		ret = max(ret, preCnt+findEnd(i, cnt, flag))
		if flag&(1<<c) == 0 {
			cnt += 1
		}
		flag |= 1 << c
	}
	return ret
}
