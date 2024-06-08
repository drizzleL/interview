package main

func printKMoves(K int) []string {
	dict := map[[2]int]rune{}
	loc := [2]int{0, 0}
	val := 'R'
	color := '_'
	dict[loc] = val
	// walk
	for i := 0; i < K; i++ {
		switch color {
		case '_':
			switch val {
			case 'U':
				val = 'R'
			case 'R':
				val = 'D'
			case 'L':
				val = 'U'
			case 'D':
				val = 'L'
			}
			color = 'X'
		case 'X':
			switch val {
			case 'U':
				val = 'L'
			case 'R':
				val = 'U'
			case 'L':
				val = 'D'
			case 'D':
				val = 'R'
			}
			color = '_'
		}
		// change old color
		dict[loc] = color

		switch val {
		case 'U':
			loc[0]--
		case 'R':
			loc[1]++
		case 'L':
			loc[1]--
		case 'D':
			loc[0]++
		}
		color = '_'
		if savedColor, ok := dict[loc]; ok {
			color = savedColor
		}
		dict[loc] = val
	}
	// render
	var top, bottom, left, right int
	for k := range dict {
		top = min(top, k[0])
		bottom = max(bottom, k[0])
		left = min(left, k[1])
		right = max(right, k[1])
	}
	data := make([][]rune, bottom-top+1)
	for i := range data {
		data[i] = make([]rune, right-left+1)
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			val, ok := dict[[2]int{i + top, j + left}]
			if ok {
				data[i][j] = val
			} else {
				data[i][j] = '_'
			}
		}
	}
	var ret []string
	for _, line := range data {
		ret = append(ret, string(line))
	}
	return ret
}
