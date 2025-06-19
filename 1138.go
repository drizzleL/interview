package main

func alphabetBoardPath(target string) string {
	board := []string{"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"}
	dict := map[byte][2]int{}
	for i := range board {
		for j := range board[i] {
			dict[board[i][j]] = [2]int{i, j}
		}
	}
	dirPath := [4]byte{'R', 'L', 'D', 'U'}
	path := func(x int, dirWeight int) []byte {
		ret := make([]byte, abs(x))
		if x < 0 {
			dirWeight += 1
		}
		for i := 0; i < abs(x); i++ {
			ret[i] = dirPath[dirWeight]
		}
		return ret
	}
	helper := func(start, end [2]int) []byte {
		var ret []byte
		if start[1] < end[1] { //
			ret = append(ret, path(end[0]-start[0], 2)...)
			ret = append(ret, path(end[1]-start[1], 0)...)

		} else { //
			ret = append(ret, path(end[1]-start[1], 0)...)
			ret = append(ret, path(end[0]-start[0], 2)...)

		}
		return ret
	}
	last := [2]int{0, 0}
	var b []byte
	for _, c := range target {
		pos := dict[byte(c)]
		b = append(b, helper(last, pos)...)
		b = append(b, '!')
		last = pos
	}
	return string(b)
}
