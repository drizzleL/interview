package main

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ret := make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	dirs := []int{-1, 0, 1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var sum, cnt int
			for _, a := range dirs {
				for _, b := range dirs {
					i2, j2 := i+a, j+b
					if i2 < 0 || j2 < 0 || i2 >= m || j2 >= n {
						continue
					}
					sum += img[i2][j2]
					cnt += 1
				}
			}
			ret[i][j] = sum / cnt
		}
	}
	return ret
}
