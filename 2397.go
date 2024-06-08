package main

func maximumRows(matrix [][]int, numSelect int) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	nums := make([]int, 0, m)
	for i := 0; i < m; i++ {
		var num int
		for j := 0; j < n; j++ {
			num <<= 1
			num |= matrix[i][j]
		}
		nums = append(nums, num)
	}
	oneCnt := func(x int) int {
		var ret int
		for x != 0 {
			ret += x & 1
			x >>= 1
		}
		return ret
	}
	var ret int
	for flag := 0; flag < 1<<n; flag++ {
		if oneCnt(flag) != n-numSelect {
			continue
		}
		var cnt int
		for _, num := range nums {
			if num&flag == 0 {
				cnt += 1
			}
		}
		ret = max(ret, cnt)
	}
	return ret
}
