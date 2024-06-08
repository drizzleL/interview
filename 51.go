package main

func solveNQueens(n int) [][]string {
	b := make([][]byte, n)
	for i := range b {
		b[i] = make([]byte, n)
		for j := range b[i] {
			b[i][j] = '.'
		}
	}
	var ret [][]string
	var helper func(row int, col int, l, r int)
	helper = func(row int, col int, l, r int) {
		if row == n {
			var tmp []string
			for i := 0; i < n; i++ {
				tmp = append(tmp, string(b[i]))
			}
			ret = append(ret, tmp)
			return
		}
		for j := 0; j < n; j++ {
			bit := 1 << j
			if col&bit != 0 || l&bit != 0 || r&bit != 0 {
				continue
			}
			b[row][j] = 'Q'
			helper(row+1, col|bit, (l|bit)<<1, (r|bit)>>1)
			b[row][j] = '.'
		}
	}
	helper(0, 0, 0, 0)
	return ret
}
