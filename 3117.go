package main

// func minimumValueSum(nums []int, andValues []int) int {
// 	var maxVal int
// 	for _, val := range andValues {
// 		maxVal = max(maxVal, val)
// 	}
// 	maxBit := bits.Len(uint(maxVal))
// 	bits := make([][]int, maxBit)
// 	for i := range bits {
// 		bits[i] = make([]int, len(nums)+1)
// 	}
// 	for i := 1; i <= len(nums); i++ {
// 		for bit := 0; bit < maxBit; bit++ {
// 			if nums[i-1]&1<<bit != 0 {
// 				bits[i][bit] += 1
// 			}
// 			bits[i][bit] += bits[i-1][bit]
// 		}
// 	}
// 	get := func(i int, val int) int {
// 		l, r := i, len(nums)
// 		for l < r {

// 		}
// 		mid := (l + r) / 2
// 		for i := 0; i < maxBit; i++ {
// 			if (1<<i)&val == 0 {

// 			}
// 		}
// 		return r
// 	}
// 	dp := make([][]int, len(nums))
// 	for i := range dp {
// 		dp[i] = make([]int, len(andValues))
// 		for j := range dp[i] {
// 			dp[i][j] = -1
// 		}
// 	}
// 	var helper func(i, j int) int
// 	helper = func(i, j int) (ret int) {
// 		if i == len(nums) && j == len(andValues) {
// 			return 0
// 		}
// 		ret = math.MaxInt32
// 		if i == len(nums) {
// 			return
// 		}
// 		if j == len(andValues) {
// 			if andValues[j-1]&nums[i] == andValues[j-1] {
// 				return helper(i+1, j) + nums[i] - nums[i-1]
// 			}
// 			return
// 		}
// 		if dp[i][j] != -1 {
// 			return dp[i][j]
// 		}
// 		defer func() {
// 			dp[i][j] = ret
// 		}()
// 		if j != 0 && nums[i]&andValues[j-1] == andValues[j-1] {
// 			ret = min(ret, helper(i+1, j)+nums[i]-nums[i-1])
// 		}
// 		idx := get(i, andValues[j])
// 		if idx < len(nums) {
// 			ret = min(ret, helper(idx+1, j+1)+nums[idx])
// 		}
// 		return
// 	}
// 	ret := helper(0, 0)
// 	if ret == math.MaxInt32 {
// 		return -1
// 	}
// 	return ret
// }
