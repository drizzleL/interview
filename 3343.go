package main

func countBalancedPermutations(num string) int {
	cnt := make([]int, 10)
	var sum int
	for _, n := range num {
		cnt[n-'0'] += 1
		sum += int(n - '0')
	}
	if sum/2*2 != sum {
		return 0
	}
	cache := map[[2]int]int{}
	var comb func(a, b int) int
	comb = func(a, b int) (ret int) {
		if b > a-b {
			return comb(a, a-b)
		}
		if c, ok := cache[[2]int{a, b}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{a, b}] = ret
		}()
		ret = 1
		for i := 0; i < b; i++ {
			ret *= (a - i)
			ret /= (i + 1)
		}
		ret %= 1e9 + 7
		return ret
	}
	size1 := (len(num) + 1) / 2
	size2 := len(num) - size1
	ans := map[[4]int]int{}
	var helper func(i int, nowCnt, cnt2, nowSum int) int
	helper = func(i int, nowCnt, anotherCnt, nowSum int) (ret int) {
		if i == 10 {
			if nowCnt == size1 && nowSum == sum/2 && anotherCnt == size2 {
				return 1
			}
			return 0
		}
		if c, ok := ans[[4]int{i, nowCnt, anotherCnt, nowSum}]; ok {
			return c
		}
		defer func() {
			ans[[4]int{i, nowCnt, anotherCnt, nowSum}] = ret
		}()
		for j := 0; j <= cnt[i] && j+nowCnt <= size1 && nowSum+j*i <= sum/2; j++ {
			if anotherCnt+cnt[i]-j > size2 {
				continue
			}
			a := comb(size1-nowCnt, j) * comb(size2-anotherCnt, cnt[i]-j)
			a %= 1e9 + 7
			a *= helper(i+1, nowCnt+j, anotherCnt+cnt[i]-j, nowSum+j*i)
			a %= 1e9 + 7
			ret += a
			ret %= 1e9 + 7
		}
		return ret
	}
	return helper(0, 0, 0, 0)
}

var combCache = map[[2]int]int{}

func comb(a, b int) (ret int) {
	if b > a-b {
		return comb(a, a-b)
	}
	if c, ok := combCache[[2]int{a, b}]; ok {
		return c
	}
	defer func() {
		combCache[[2]int{a, b}] = ret
	}()
	ret = 1
	for i := 0; i < b; i++ {
		ret *= a - i
		ret %= 1e9 + 7
	}
	for i := 0; i < b; i++ {
		ret /= i + 1
	}
	ret %= 1e9 + 7
	return ret
}
