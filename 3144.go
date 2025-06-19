package main

func minimumSubstringsInPartition(s string) int {
	cache := make([]int, len(s))
	for i := range cache {
		cache[i] = -1
	}
	check := func(dict [26]int) bool {
		var flag int
		for _, v := range dict {
			if v == 0 || v == flag {
				continue
			}
			if flag != 0 {
				return false
			}
			flag = v
		}
		return true
	}
	var helper func(i int) int
	helper = func(i int) (ret int) {
		if i == len(s) {
			return 0
		}
		if c := cache[i]; c != -1 {
			return c
		}
		defer func() {
			cache[i] = ret
		}()
		ret = len(s) - i
		dict := [26]int{}
		for j := i; j < len(s); j++ {
			dict[s[j]-'a'] += 1
			if !check(dict) {
				continue
			}
			ret = min(ret, 1+helper(j+1))
		}
		return
	}
	return helper(0)
}
