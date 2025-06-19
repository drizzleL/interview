package main

func rotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])
	ret := make([][]byte, n)
	for i := range ret {
		ret[i] = make([]byte, m)
	}
	helper := func(i, j int, cnt int) {
		for a, b := j, m-1-i; cnt != 0; a, cnt = a-1, cnt-1 {
			ret[a][b] = '#'
		}
	}
	for i := 0; i < m; i++ {
		var cnt int
		for j := 0; j < n; j++ {
			ret[j][m-1-i] = '.'
			switch box[i][j] {
			case '.':
				continue
			case '#':
				cnt += 1
			case '*':
				helper(i, j-1, cnt)
				cnt = 0
				ret[j][m-1-i] = '*'
			}
		}
		helper(i, n-1, cnt)
	}
	return ret
}
