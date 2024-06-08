package main

func tictactoe(board []string) string {
	m := len(board)
	var flag bool
	for i := 0; i < m; i++ {
		var xcnt, ocnt int
		for j := 0; j < m; j++ {
			switch board[i][j] {
			case ' ':
				flag = true
			case 'X':
				xcnt++
			case 'O':
				ocnt++
			}
		}
		if xcnt == m {
			return "X"
		}
		if ocnt == m {
			return "O"
		}
	}
	for j := 0; j < m; j++ {
		var xcnt, ocnt int
		for i := 0; i < m; i++ {
			switch board[i][j] {
			case 'X':
				xcnt++
			case 'O':
				ocnt++
			}
		}
		if xcnt == m {
			return "X"
		}
		if ocnt == m {
			return "O"
		}
	}
	{
		var xcnt, ocnt int
		for i := 0; i < m; i++ {
			switch board[i][i] {
			case 'X':
				xcnt++
			case 'O':
				ocnt++
			}
		}
		if xcnt == m {
			return "X"
		}
		if ocnt == m {
			return "O"
		}
	}
	{
		var xcnt, ocnt int
		for i := m - 1; i >= 0; i-- {
			switch board[i][m-i-1] {
			case 'X':
				xcnt++
			case 'O':
				ocnt++
			}
		}
		if xcnt == m {
			return "X"
		}
		if ocnt == m {
			return "O"
		}
	}
	if flag {
		return "Pending"
	}
	return "Draw"
}
