package main

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	var ret int
	heights := make([]int, n)
	var stack []int
	for r := 0; r < m; r++ {
		for c := 0; c <= n; c++ {
			if c != n {
				if matrix[r][c] == '0' {
					heights[c] = 0
				} else {
					heights[c] += 1
				}
			}
			for len(stack) != 0 && (c == n || heights[c] < heights[stack[len(stack)-1]]) {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				width := c
				if len(stack) != 0 {
					width = c - stack[len(stack)-1] - 1
				}
				ret = max(ret, width*heights[top])

			}
			if c != n {
				stack = append(stack, c)
			}
		}
	}
	return ret
}
