package main

func queryResults(limit int, queries [][]int) []int {
	colors := map[int]int{}
	colorCnt := map[int]int{}
	ret := make([]int, 0, len(queries))
	var tmp int
	for _, q := range queries {
		num, col := q[0], q[1]
		oldColor := colors[num]
		if oldColor != 0 {
			colorCnt[oldColor] -= 1
			if colorCnt[oldColor] == 0 {
				tmp -= 1
			}
		}
		colors[num] = col
		colorCnt[col] += 1
		if colorCnt[col] == 1 {
			tmp += 1
		}
		ret = append(ret, tmp)
	}
	return ret
}
