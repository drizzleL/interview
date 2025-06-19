package main

import "math"

func mctFromLeafValues2(arr []int) int {
	var stack []int
	stack = append(stack, math.MaxInt32)
	var ret int
	for _, num := range arr {
		for stack[len(stack)-1] <= num {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret += top * min(num, stack[len(stack)-1])
		}
		stack = append(stack, num)
	}
	for len(stack) > 2 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret += top * stack[len(stack)-1]
	}
	return ret
}
func mctFromLeafValues(arr []int) int {
	dp := make([][]int, len(arr))
	maxVal := make([][]int, len(arr))
	var sum int
	for i := range dp {
		dp[i] = make([]int, len(arr))
		dp[i][i] = arr[i]
		maxVal[i] = make([]int, len(arr))
		maxVal[i][i] = arr[i]
		for j := i + 1; j < len(arr); j++ {
			dp[i][j] = math.MaxInt32
			maxVal[i][j] = max(maxVal[i][j-1], arr[j])
		}
		sum += arr[i]
	}
	for length := 2; length <= len(arr); length++ {
		for i := 0; i+length-1 < len(arr); i++ {
			j := i + length - 1
			for m := i; m < j; m++ {
				dp[i][j] = min(dp[i][j], dp[i][m]+dp[m+1][j]+maxVal[i][m]*maxVal[m+1][j])
			}
		}
	}
	return dp[0][len(arr)-1] - sum
}
