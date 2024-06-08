package main

import "math"

func getMaxMatrix(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	area := make([][]int, m)
	for i := range area {
		area[i] = make([]int, n)
	}
	maxSum := math.MinInt32
	var ret []int

	for left := 0; left < n; left++ {
		tmpSum := make([]int, m)
		for right := left; right < n; right++ {
			for i := 0; i < m; i++ {
				tmpSum[i] += matrix[i][right]
			}
			var maxTop, maxBottom int
			sum := tmpSum[0]
			var currTop, currBottom int
			currSum := tmpSum[0]

			for i := 1; i < m; i++ {
				if tmpSum[i] > tmpSum[i]+currSum {
					currSum = tmpSum[i]
					currTop = i
					currBottom = i
				} else {
					currSum += tmpSum[i]
					currBottom = i
				}
				if currSum > sum {
					sum = currSum
					maxTop, maxBottom = currTop, currBottom
				}
			}
			if sum > maxSum {
				maxSum = sum
				ret = []int{maxTop, left, maxBottom, right}
			}
		}
	}
	return ret
}
