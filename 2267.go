package main

func hasValidPath2(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	if grid[0][0] == ')' {
		return false
	}
	if grid[m-1][n-1] == '(' {
		return false
	}
	if (m+n-1)%2 != 0 {
		return false
	}
	helper := func(last map[int]bool, now map[int]bool, cnt, flag int) bool {
		var ret bool
		for k := range last {
			right := k + flag
			left := cnt - right
			if left < right {
				continue
			}
			if left > (m+n-1)/2 {
				continue
			}
			now[right] = true
			ret = true
		}
		return ret
	}
	last := make([]map[int]bool, n)
	for j := 0; j < n; j++ {
		last[j] = make(map[int]bool)
		if j == 0 {
			last[j][0] = true
			continue
		}
		helper(last[j-1], last[j], j+1, int(grid[0][j]-'('))
	}
	for i := 1; i < m; i++ {
		now := make([]map[int]bool, n)
		var found bool
		for j := 0; j < n; j++ {
			cnt := i + j + 1
			flag := int(grid[i][j] - '(')
			now[j] = make(map[int]bool)
			if j != 0 {
				ok := helper(now[j-1], now[j], cnt, flag)
				found = found || ok
			}
			if i != 0 {
				ok := helper(last[j], now[j], cnt, flag)
				found = found || ok
			}
		}
		if !found {
			return false
		}
		last = now
	}
	return len(last[n-1]) != 0
}
