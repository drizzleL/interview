package main

func countCoprime(mat [][]int) int {
	dp := make([]int, 151)
	for j := 0; j < len(mat[0]); j++ {
		dp[mat[0][j]] += 1
	}
	for i := 1; i < len(mat); i++ {
		newDp := make([]int, 151)
		for j := 0; j < len(mat[i]); j++ {
			val := mat[i][j]
			for idx := 1; idx <= 150; idx++ {
				g := gcd(val, idx)
				newDp[g] += dp[idx]
			}
		}
		for i := range newDp {
			newDp[i] %= 1e9 + 7
		}
		dp = newDp
	}
	return dp[1]
}
