package main

func solveSudoku(board [][]byte) {
	var rows, cols, blocks [9]int
	var q [][2]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				q = append(q, [2]int{i, j})
				continue
			}
			rows[i] |= 1 << (board[i][j] - '1')
			cols[j] |= 1 << (board[i][j] - '1')
			blocks[i/3*3+j/3] |= 1 << (board[i][j] - '1')
		}
	}
	var solve func(q [][2]int) bool
	solve = func(q [][2]int) bool {
		if len(q) == 0 {
			return true
		}
		i, j := q[len(q)-1][0], q[len(q)-1][1]
		blockNum := i/3*3 + j/3
		row, col, block := rows[i], cols[j], blocks[blockNum]
		mask := row | col | block
		for v := 0; v < 9; v++ {
			if mask&(1<<v) != 0 {
				continue
			}
			board[i][j] = '1' + byte(v)
			rows[i] |= 1 << v
			cols[j] |= 1 << v
			blocks[blockNum] |= 1 << v
			if solve(q[:len(q)-1]) {
				return true
			}
			rows[i] ^= 1 << v
			cols[j] ^= 1 << v
			blocks[blockNum] ^= 1 << v
			board[i][j] = '.'
		}
		rows[i], cols[j], blocks[blockNum] = row, col, block
		return false
	}
	solve(q)
}
