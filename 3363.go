package main

func maxCollectedFruits(fruits [][]int) int {
	n := len(fruits)
	var sum int
	for i := 0; i < n; i++ {
		sum += fruits[i][i]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fruits[i][j] = 0
		}
	}
	for i := 0; i < n; i++ {
		fruits[i][i] = 0
	}
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			tmp := max(fruits[i-1][j-1], fruits[i-1][j])
			if j+1 < n {
				tmp = max(tmp, fruits[i-1][j+1])
			}
			fruits[i][j] += tmp
		}
	}
	sum += fruits[n-2][n-1]
	for j := 1; j < n-1; j++ {
		for i := j + 1; i < n; i++ {
			tmp := max(fruits[i-1][j-1], fruits[i][j-1])
			if i+1 < n {
				tmp = max(tmp, fruits[i+1][j-1])
			}
			fruits[i][j] += tmp
		}
	}
	sum += fruits[n-1][n-2]
	return sum
}
