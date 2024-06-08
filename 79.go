package main

func exist(board [][]byte, word string) bool {
	seen := make([][]bool, len(board))
	for i := range seen {
		seen[i] = make([]bool, len(board[0]))
	}
	var helper func(i int, now [2]int) bool
	helper = func(i int, now [2]int) bool {
		if i == len(word) {
			return true
		}
		if board[now[0]][now[1]] != word[i] {
			return false
		}
		if i == len(word)-1 {
			return true
		}
		seen[now[0]][now[1]] = true
		for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			i2, j2 := now[0]+dir[0], now[1]+dir[1]
			if i2 < 0 || j2 < 0 || i2 >= len(board) || j2 >= len(board[0]) {
				continue
			}
			if seen[i2][j2] {
				continue
			}
			if helper(i+1, [2]int{i2, j2}) {
				return true
			}
		}
		seen[now[0]][now[1]] = false
		return false
	}
	for i := range board {
		for j := range board[i] {
			if helper(0, [2]int{i, j}) {
				return true
			}
		}
	}
	return false
}
