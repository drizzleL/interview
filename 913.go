package main

func catMouseGame(graph [][]int) int {
	size := len(graph)
	dp := make([][][]int, size)
	degree := make([][][]int, size)
	for i := range dp {
		dp[i] = make([][]int, size)
		degree[i] = make([][]int, size)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			degree[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			degree[i][j][0] = len(graph[i])
			degree[i][j][1] = len(graph[j])
			for _, v := range graph[j] {
				if v == 0 {
					degree[i][j][1] -= 1
					break
				}
			}
		}
	}
	var q [][3]int
	for j := 1; j < size; j++ {
		for k := 0; k < 2; k++ {
			dp[0][j][k] = 1
			q = append(q, [3]int{0, j, k})
			dp[j][j][k] = 2
			q = append(q, [3]int{j, j, k})
		}
	}
	for len(q) > 0 {
		top := q[len(q)-1]
		q = q[:len(q)-1]
		mouse, cat, turn := top[0], top[1], top[2]
		ret := dp[mouse][cat][turn]
		if mouse == 1 && cat == 2 && turn == 0 {
			return dp[1][2][0]
		}
		prevTurn := 1 - turn
		x := mouse
		if prevTurn == 1 {
			x = cat
		}
		for _, prev := range graph[x] {
			prevCat, prevMouse := cat, mouse
			if prevTurn == 0 {
				prevMouse = prev
			} else {
				prevCat = prev
			}
			if prevCat == 0 {
				continue
			}
			if dp[prevMouse][prevCat][prevTurn] != 0 {
				continue
			}
			degree[prevMouse][prevCat][prevTurn] -= 1
			if prevTurn == 0 && ret == 1 || prevTurn == 1 && ret == 2 || degree[prevMouse][prevCat][prevTurn] == 0 {
				dp[prevMouse][prevCat][prevTurn] = ret
				q = append(q, [3]int{prevMouse, prevCat, prevTurn})
			}
		}
	}
	return 0
}
