package main

func executeInstructions(n int, startPos []int, s string) []int {
	ret := make([]int, len(s))
	var curr [2]int
	dirs := map[byte][2]int{
		'R': {0, 1},
		'L': {0, -1},
		'U': {-1, 0},
		'D': {1, 0},
	}
	upmost := startPos[0] + 1
	downmost := n - startPos[0]
	leftmost := startPos[1] + 1
	rightmost := n - startPos[1]
	nextRow, nextCol := map[int]int{}, map[int]int{}
	nextRow[0] = len(s)
	nextCol[0] = len(s)
	for i := len(s) - 1; i >= 0; i-- {
		curr[0] -= dirs[s[i]][0]
		curr[1] -= dirs[s[i]][1]
		tmp := len(s) - i
		if _, ok := nextRow[curr[0]-upmost]; ok {
			tmp = min(tmp, nextRow[curr[0]-upmost]-i-1)
		}
		if _, ok := nextRow[curr[0]+downmost]; ok {
			tmp = min(tmp, nextRow[curr[0]+downmost]-i-1)
		}
		if _, ok := nextCol[curr[1]-leftmost]; ok {
			tmp = min(tmp, nextCol[curr[1]-leftmost]-i-1)
		}
		if _, ok := nextCol[curr[1]+rightmost]; ok {
			tmp = min(tmp, nextCol[curr[1]+rightmost]-i-1)
		}
		nextRow[curr[0]] = i
		nextCol[curr[1]] = i
		ret[i] = tmp
	}
	return ret
}
