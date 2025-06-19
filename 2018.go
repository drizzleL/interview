package main

func placeWordInCrossword(board [][]byte, word string) bool {
	checkHor := func(i int, start, end int) bool {
		if end-start+1 != len(word) {
			return false
		}
		var flag bool
		for j, m := start, 0; j <= end; j, m = j+1, m+1 {
			if board[i][j] != ' ' && word[m] != board[i][j] {
				flag = true
				break
			}
		}
		if !flag {
			return true
		}
		flag = false
		for j, m := start, len(word)-1; j <= end; j, m = j+1, m-1 {
			if board[i][j] != ' ' && word[m] != board[i][j] {
				flag = true
				break
			}
		}
		return !flag
	}
	checkVer := func(j int, start, end int) bool {
		if end-start+1 != len(word) {
			return false
		}
		var flag bool
		for i, m := start, 0; i <= end; i, m = i+1, m+1 {
			if board[i][j] != ' ' && word[m] != board[i][j] {
				flag = true
				break
			}
		}
		if !flag {
			return true
		}
		flag = false
		for i, m := start, len(word)-1; i <= end; i, m = i+1, m-1 {
			if board[i][j] != ' ' && word[m] != board[i][j] {
				flag = true
				break
			}
		}
		if !flag {
			return true
		}
		return true
	}
	for i := 0; i < len(board); i++ {
		var start int
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '#' {
				if checkHor(i, start, j-1) {
					return true
				}
				start = j + 1
			}
		}
		if checkHor(i, start, len(board[0])-1) {
			return true
		}
	}
	for j := 0; j < len(board[0]); j++ {
		var start int
		for i := 0; i < len(board); i++ {
			if board[i][j] == '#' {
				if checkVer(j, start, i-1) {
					return true
				}
				start = i + 1
			}
		}
		if checkVer(j, start, len(board)-1) {
			return true
		}
	}
	return false
}
