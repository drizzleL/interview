package main

import (
	"math"
	"math/bits"
	"sort"
)

func minimumDifference3(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	uniq := func(nums []int) []int {
		i := 1
		for j := 1; j < len(nums); j++ {
			if nums[j] == nums[j-1] {
				continue
			}
			nums[i] = nums[j]
			i++
		}
		return nums[:i]
	}
	n := len(nums) / 2
	helper := func(nums []int) [][]int {
		dp := make([][]int, n+1)
		for mask := 0; mask < 1<<n; mask++ {
			take := bits.OnesCount(uint(mask))
			var tmp int
			for j := range nums {
				if mask&(1<<j) == 0 {
					continue
				}
				tmp += nums[j]
			}
			dp[take] = append(dp[take], tmp)
		}
		for i := range dp {
			sort.Ints(dp[i])
			dp[i] = uniq(dp[i])
		}
		return dp
	}
	s1, s2 := helper(nums[:n]), helper(nums[n:])
	ret := math.MaxInt32
	f := func(sum1 int) int {
		sum2 := sum - sum1
		return abs(sum1 - sum2)
	}
	for k := 0; k <= n; k++ {
		for i, j := 0, len(s2[n-k])-1; i < len(s1[k]); i++ {
			expected := sum/2 - s1[k][i]
			for j > 0 && s2[n-k][j-1] >= expected {
				j--
			}
			ret = min(ret, f(s1[k][i]+s2[n-k][j]))
			if j > 0 {
				ret = min(ret, f(s1[k][i]+s2[n-k][j-1]))
			}
		}
	}
	return ret
}
