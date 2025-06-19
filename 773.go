package main

func slidingPuzzle(board [][]int) int {
	var v [6]int
	for i, k := 0, 0; i <= 1; i++ {
		for j := 0; j <= 2; j++ {
			v[k] = board[i][j]
			k++
		}
	}
	q := [][6]int{v}
	seen := map[[6]int]bool{
		v: true,
	}
	dest := [6]int{1, 2, 3, 4, 5, 0}
	if dest == v {
		return 0
	}
	check := func(ele [6]int, next *[][6]int) bool {
		if dest == ele {
			return true
		}
		if seen[ele] {
			return false
		}
		seen[ele] = true
		*next = append(*next, ele)
		return false
	}
	for step := 1; len(q) != 0; step++ {
		var next [][6]int
		for _, ele := range q {
			var idx int
			for i, v := range ele {
				if v == 0 {
					idx = i
					break
				}
			}
			switch idx {
			case 2, 5:
			default:
				v2 := ele
				v2[idx+1], v2[idx] = v2[idx], v2[idx+1]
				if check(v2, &next) {
					return step
				}
			}
			switch idx {
			case 0, 3:
			default:
				v2 := ele
				v2[idx-1], v2[idx] = v2[idx], v2[idx-1]
				if check(v2, &next) {
					return step
				}
			}
			switch idx {
			case 0, 1, 2:
				v2 := ele
				v2[idx+3], v2[idx] = v2[idx], v2[idx+3]
				if check(v2, &next) {
					return step
				}
			}
			switch idx {
			case 3, 4, 5:
				v2 := ele
				v2[idx-3], v2[idx] = v2[idx], v2[idx-3]
				if check(v2, &next) {
					return step
				}
			}
		}
		q = next
	}
	return -1
}
