package main

func judgeCircle(moves string) bool {
	pos := [2]int{0, 0}
	for _, m := range moves {
		switch m {
		case 'U':
			pos[0]++
		case 'D':
			pos[0]--
		case 'L':
			pos[1]++
		case 'R':
			pos[1]--
		}
	}
	return pos == [2]int{0, 0}
}
