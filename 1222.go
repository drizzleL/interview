package main

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	seen := make([]bool, 64)
	for _, q := range queens {
		i, j := q[0], q[1]
		seen[i*8+j] = true
	}
	var ret [][]int
	dirs := []int{-1, 0, 1}
	for _, d1 := range dirs {
		for _, d2 := range dirs {
			if d1 == 0 && d2 == 0 {
				continue
			}
			i, j := king[0]+d1, king[1]+d2
			for i >= 0 && j >= 0 && i < 8 && j < 8 {
				if seen[i*8+j] {
					ret = append(ret, []int{i, j})
					break
				}
				i += d1
				j += d2
			}
		}
	}
	return ret
}
