package main

func rotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ret := make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	var next func(x, y int, k int, up, bottom, left, right int) (x2, y2 int)
	next = func(x, y int, k int, up, bottom, left, right int) (x2, y2 int) {
		if x == up { // go left
			if y-left >= k {
				return x, y - k
			}
			return next(x+1, left, k-(y-left)-1, up, bottom, left, right)
		}
		if x == bottom { // go right
			if right-y >= k {
				return x, y + k
			}
			return next(x-1, right, k-(right-y)-1, up, bottom, left, right)
		}
		if y == left { // go down
			if bottom-x >= k {
				return x + k, y
			}
			return next(bottom, y+1, k-(bottom-x)-1, up, bottom, left, right)
		}
		if y == right { // go up
			if x-up >= k {
				return x - k, y
			}
			return next(up, y-1, k-(x-up)-1, up, bottom, left, right)
		}
		return 0, 0
	}
	for i, j := 0, 0; i*2 < m && j*2 < n; i, j = i+1, j+1 {
		s := (m-i*2-1)*2 + (n-j*2-1)*2
		kk := k % s
		left, right := j, n-j-1
		up, bottom := i, m-i-1
		for i2 := up; i2 <= bottom; i2++ {
			var x2, y2 int
			x2, y2 = next(i2, left, kk, up, bottom, left, right)
			ret[x2][y2] = grid[i2][left]
			x2, y2 = next(i2, right, kk, up, bottom, left, right)
			ret[x2][y2] = grid[i2][right]
		}
		for j2 := left; j2 <= right; j2++ {
			var x2, y2 int
			x2, y2 = next(up, j2, kk, up, bottom, left, right)
			ret[x2][y2] = grid[up][j2]
			x2, y2 = next(bottom, j2, kk, up, bottom, left, right)
			ret[x2][y2] = grid[bottom][j2]
		}
	}
	return ret
}
