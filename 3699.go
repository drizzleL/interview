package main

func zigZagArrays(n int, l int, r int) int {
	inc, dec := make([][]int, n+1), make([][]int, n+1)
	for i := range inc {
		inc[i] = make([]int, r+2)
		dec[i] = make([]int, r+2)
	}
	inc[0][r+1] = 1
	dec[0][l-1] = 1
	for i := 1; i <= n; i++ {
		for j := l; j <= r; j++ {
			inc[i][j] += inc[i-1][j-1] + dec[i-1][j-1]
			inc[i][j] %= 1e9 + 7
		}
		for j := r; j >= l; j-- {
			dec[i][j] += dec[i-1][j+1] + inc[i-1][j+1]
			dec[i][j] %= 1e9 + 7
		}
	}
	var ret int
	for i := l; i <= r; i++ {
		ret += inc[n][i] + dec[n][i]
		ret %= 1e9 + 7
	}
	return ret
}
