package main

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	cache := map[[3]int][]int{}
	var dfs func(l, r, m int) []int
	dfs = func(l, r, m int) (ret []int) {
		if l == r {
			return []int{1, 1}
		}
		if l > r {
			return dfs(r, l, m)
		}
		if c, ok := cache[[3]int{l, r, m}]; ok {
			return c
		}
		defer func() {
			cache[[3]int{l, r, m}] = ret
		}()
		ret = []int{n, 0}
		for i := 0; i <= l; i++ {
			nextM := (m + 1) / 2
			_, lLost := i, l-i
			if r < m/2 {
				for j := lLost; j <= r-l-1+lLost; j++ {
					tmp := dfs(i, j, nextM)
					ret[0] = min(ret[0], tmp[0])
					ret[1] = max(ret[1], tmp[1])
				}
			} else {
				extra := nextM - (m - r)
				for j := extra + lLost; j <= extra+m-l-r-2+lLost; j++ {
					tmp := dfs(i, j, nextM)
					ret[0] = min(ret[0], tmp[0])
					ret[1] = max(ret[1], tmp[1])
				}
			}
		}
		ret[0] += 1
		ret[1] += 1
		return
	}
	return dfs(firstPlayer-1, n-secondPlayer, n)
}
