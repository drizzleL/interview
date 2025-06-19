package main

func minOperations20(word1 string, word2 string) int {
	cache := map[[2]int]int{}
	helper2 := func(i, j int, rev bool) int {
		var ret int
		if rev {
			ret = 1
		}
		dict := map[[2]byte]int{}
		for m := i; m <= j; m++ {
			m2 := m
			if rev {
				m2 = j - m + i
			}
			if word1[m] == word2[m2] {
				continue
			}
			k := [2]byte{word1[m], word2[m2]}
			if dict[k] != 0 {
				dict[k] -= 1
				ret += 1
				continue
			}
			k2 := [2]byte{word2[m2], word1[m]}
			dict[k2] += 1
		}
		for _, v := range dict {
			ret += v
		}
		return ret
	}
	var helper func(i, j int) int
	helper = func(i, j int) (ret int) {
		if c, ok := cache[[2]int{i, j}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{i, j}] = ret
		}()
		ret = min(helper2(i, j, false), helper2(i, j, true))
		if ret <= 1 {
			return
		}
		for mid := i; mid < j; mid++ {
			ret = min(ret, helper(i, mid)+helper(mid+1, j))
		}
		return
	}
	return helper(0, len(word1)-1)
}
