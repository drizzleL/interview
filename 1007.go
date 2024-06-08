package main

func minDominoRotations(tops []int, bottoms []int) int {
	size := len(tops)
	ret := size

	points := make([][]int, 2)
	points[0] = tops
	points[1] = bottoms

	helper := func(v int, idx int) int {
		var ret int
		sideIdx := 1 - idx
		for i := 0; i < size; i++ {
			if points[idx][i] == v {
				continue
			}
			if points[sideIdx][i] != v {
				return size
			}
			ret++
		}
		return ret
	}
	for i := 1; i <= 6; i++ {
		ret = min(ret, helper(i, 0))
		ret = min(ret, helper(i, 1))
	}

	if ret == size {
		return -1
	}
	return ret
}
