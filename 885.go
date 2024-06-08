package main

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	total := rows * cols
	var cnt int
	var ret [][]int
	node := []int{rStart, cStart}
	step := 1
	check := func() {
		if node[0] >= 0 && node[0] < rows && node[1] >= 0 && node[1] < cols {
			ret = append(ret, []int{node[0], node[1]})
			cnt += 1
		}
	}
	check()
	for cnt < total {
		for i := 0; i < step; i++ {
			node[1] += 1
			check()
		}
		for i := 0; i < step; i++ {
			node[0] += 1
			check()
		}
		step += 1
		for i := 0; i < step; i++ {
			node[1] -= 1
			check()
		}
		for i := 0; i < step; i++ {
			node[0] -= 1
			check()
		}
		step += 1
	}
	return ret
}
