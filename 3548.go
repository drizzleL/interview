package main

func canPartitionGrid2(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	dict1 := map[int]int{}
	dict2 := map[int]int{}
	dict3 := map[int]int{}
	dict4 := map[int]int{}
	var sum int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum += grid[i][j]
			dict2[grid[i][j]] += 1
			dict4[grid[i][j]] += 1
		}
	}
	var sum1, sum2 int
	for i := 0; i < m-1; i++ {
		for j := 0; j < n; j++ {
			sum1 += grid[i][j]
			dict1[grid[i][j]] += 1
			dict2[grid[i][j]] -= 1
		}
		if sum1*2 == sum {
			return true
		}
		if sum1*2 > sum && dict1[sum1*2-sum] > 0 {
			if n == 1 {
				if i > 0 && (grid[0][0] == sum1*2-sum || grid[i][0] == sum1*2-sum) {
					return true
				}
			} else {
				if i > 0 || (grid[0][0] == sum1*2-sum || grid[0][n-1] == sum1*2-sum) {
					return true
				}
			}
		}
		if sum1*2 < sum && dict2[sum-sum1*2] > 0 {
			if n == 1 {
				if i != m-2 && (grid[m-1][0] == sum-sum1*2 || grid[i+1][0] == sum-sum1*2) {
					return true
				}
			} else {
				if i != m-2 || (grid[m-1][0] == sum-sum1*2 || grid[m-1][n-1] == sum-sum1*2) {
					return true
				}
			}
		}
	}

	for j := 0; j < n-1; j++ {
		for i := 0; i < m; i++ {
			sum2 += grid[i][j]
			dict3[grid[i][j]] += 1
			dict4[grid[i][j]] -= 1
		}
		if sum2*2 == sum {
			return true
		}
		if sum2*2 > sum && dict3[sum2*2-sum] > 0 {
			if m == 1 {
				if j > 0 && (grid[0][0] == sum2*2-sum || grid[0][j] == sum2*2-sum) {
					return true
				}
			} else {
				if j > 0 || (grid[0][0] == sum2*2-sum || grid[m-1][0] == sum2*2-sum) {
					return true
				}
			}
		}
		if sum2*2 < sum && dict4[sum-sum2*2] > 0 {
			if m == 1 {
				if j != n-2 && (grid[0][n-1] == sum-sum2*2 || grid[0][j+1] == sum-sum2*2) {
					return true
				}
			} else {
				if j != n-2 || (grid[0][n-1] == sum-sum2*2 || grid[m-1][n-1] == sum-sum2*2) {
					return true
				}
			}
		}
	}
	return false
}
