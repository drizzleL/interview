package main

func mostFrequentPrime(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	dict := map[int]int{}
	isPrime := func(x int) bool {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	var maxVal int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := mat[i][j]
			for _, dir := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
				i2, j2, v2 := i, j, v
				for {
					i2, j2 = i2+dir[0], j2+dir[1]
					if i2 < 0 || j2 < 0 || i2 >= m || j2 >= n {
						break
					}
					v2 = v2*10 + mat[i2][j2]
					if !isPrime(v2) {
						continue
					}
					dict[v2] += 1
					maxVal = max(maxVal, v2)
				}
			}
		}
	}
	var maxCnt int
	ret := -1
	for k, v := range dict {
		if v < maxCnt {
			continue
		}
		if v == maxCnt && k < ret {
			continue
		}
		ret = k
		maxCnt = v
	}
	return ret
}
