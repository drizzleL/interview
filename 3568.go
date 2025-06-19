package main

func minMoves4(classroom []string, energy int) int {
	m, n := len(classroom), len(classroom[0])
	var start int
	var ll []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch classroom[i][j] {
			case 'S':
				start = i*n + j
			case 'L':
				ll = append(ll, i*n+j)
			}
		}
	}
	literDict := map[int]int{}
	for i, l := range ll {
		literDict[l] = i
	}
	dp := make([][]int, m*n)
	for i := range dp {
		dp[i] = make([]int, 1<<len(ll))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	type node struct {
		idx       int
		literMask int
		energy    int
		move      int
	}
	endMask := (1 << len(ll)) - 1
	var q []node
	q = append(q, node{idx: start, literMask: 0, energy: energy, move: 0})
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		if top.literMask == endMask {
			return top.move
		}
		x, y := top.idx/n, top.idx%n
		for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			x2, y2 := x+d[0], y+d[1]
			if x2 < 0 || x2 >= m || y2 < 0 || y2 >= n {
				continue
			}
			if classroom[x2][y2] == 'X' {
				continue
			}
			en := top.energy
			en -= 1
			if en < 0 {
				continue
			}
			if classroom[x2][y2] == 'R' {
				en = energy
			}
			if dp[x2*n+y2][top.literMask] >= en {
				continue
			}
			dp[x2*n+y2][top.literMask] = en
			mask := top.literMask
			if classroom[x2][y2] == 'L' {
				mask |= 1 << literDict[x2*n+y2]
			}
			q = append(q, node{
				idx:       x2*n + y2,
				literMask: mask,
				move:      top.move + 1,
				energy:    en,
			})
		}
	}
	return -1
}
