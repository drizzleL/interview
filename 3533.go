package main

import (
	"math"
	"sort"
)

func concatenatedDivisibility(nums []int, k int) []int {
	sort.Ints(nums)
	if k == 1 {
		return nums
	}
	size := 1 << len(nums)
	dp := make([][][]int, size)
	for i := range dp {
		dp[i] = make([][]int, k)
		for j := range dp[i] {
			dp[i][j] = []int{-1}
		}
	}
	check := func(a []int) bool {
		return !(len(a) == 1 && a[0] == -1)
	}
	helper := func(a, b []int) []int {
		ret := make([]int, len(b))
		if !check(a) {
			copy(ret, b)
			return b
		}
		var ans []int
		for i := range a {
			if a[i] == b[i] {
				continue
			}
			if a[i] < b[i] {
				ans = a
			} else {
				ans = b
			}
			break
		}
		copy(ret, ans)
		return ret
	}
	dp[0][0] = []int{}
	for mask := 1; mask < size; mask++ {
		for i := 0; i < len(nums); i++ {
			if (mask>>i)&1 == 0 { // can't reach from here
				continue
			}
			before := mask ^ (1 << i)
			for rem := 0; rem < k; rem++ {
				if !check(dp[before][rem]) {
					continue
				}
				newRem := (rem*int(math.Pow10(1+int(math.Log10(float64(nums[i]))))) + nums[i]) % k
				dp[mask][newRem] = helper(dp[mask][newRem], append(dp[before][rem], i))
			}
		}
	}
	if !check(dp[size-1][0]) {
		return nil
	}
	var ret []int
	for _, i := range dp[size-1][0] {
		ret = append(ret, nums[i])
	}
	return ret
}
