package main

func stoneGameII(piles []int) int {
	presum := make([]int, len(piles)+1)
	for i := range piles {
		presum[i+1] = presum[i] + piles[i]
	}
	getSum := func(i, j int) int {
		return presum[j+1] - presum[i]
	}
	cache := make([][]int, len(piles))
	for i := range cache {
		cache[i] = make([]int, len(piles))
	}
	var helper func(i, m int) int
	helper = func(i, m int) (ret int) {
		if len(piles)-i <= m*2 { // takes all
			return getSum(i, len(piles)-1)
		}
		if cache[i][m] > 0 {
			return cache[i][m]
		}
		defer func() {
			cache[i][m] = ret
		}()
		for k := 1; k <= m*2; k++ {
			ret = max(ret, getSum(i, len(piles)-1)-helper(i+k, max(k, m)))
		}
		return
	}
	return helper(0, 1)
}
