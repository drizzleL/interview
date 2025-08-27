package main

import "math"

func minStable(nums []int, maxC int) int {
	var oneCnt int
	for _, num := range nums {
		if num == 1 {
			oneCnt += 1
		}
	}
	if maxC >= len(nums)-oneCnt {
		return 0
	}
	dp := make([][]int, len(nums))
	size := int(math.Log2(float64(len(nums))))
	for i := range dp {
		dp[i] = make([]int, size+1)
	}
	for i := 0; i < len(nums); i++ {
		dp[i][0] = nums[i]
	}
	for bit := 1; bit <= size; bit++ {
		for i := 0; i < len(nums); i++ {
			if i+(1<<bit)-1 >= len(nums) {
				break
			}
			dp[i][bit] = gcd(dp[i][bit-1], dp[i+(1<<(bit-1))][bit-1])
		}
	}
	var query func(l, r int) int
	query = func(l, r int) int {
		if l == r {
			return dp[l][0]
		}
		gap := r - l + 1
		idx := int(math.Log2(float64(gap)))
		return gcd(dp[l][idx], dp[r-(1<<idx)+1][idx])
	}
	check := func(x int) bool {
		var cnt int
		for i := 0; i < len(nums) && cnt <= maxC; i++ {
			r := i + x
			if r >= len(nums) {
				break
			}
			v := query(i, r)
			if v == 1 {
				continue
			}
			cnt += 1
			i = r
		}
		return cnt <= maxC
	}
	l, r := 1, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
