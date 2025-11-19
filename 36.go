package main

func isValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		var dict [10]bool
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			d := board[row][col] - '0'
			if dict[d] {
				return false
			}
			dict[d] = true
		}
	}
	for col := 0; col < 9; col++ {
		var dict [10]bool
		for row := 0; row < 9; row++ {
			if board[row][col] == '.' {
				continue
			}
			d := board[row][col] - '0'
			if dict[d] {
				return false
			}
			dict[d] = true
		}
	}
	for startI := 0; startI < 9; startI += 3 {
		for startJ := 0; startJ < 9; startJ += 3 {
			var dict [10]bool
			for i := startI; i < startI+3; i++ {
				for j := startJ; j < startJ+3; j++ {
					if board[i][j] == '.' {
						continue
					}
					d := board[i][j] - '0'
					if dict[d] {
						return false
					}
					dict[d] = true
				}
			}
		}
	}
	return true
}
