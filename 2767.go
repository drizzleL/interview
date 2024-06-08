package main

func minimumBeautifulSubstrings(s string) int {
	var helper func(i int) int

	var powerOfFive func(v int) bool
	powerOfFive = func(v int) bool {
		if v == 1 {
			return true
		}
		if v%5 != 0 {
			return false
		}
		return powerOfFive(v / 5)
	}
	cache := make([]int, len(s))
	helper = func(i int) (ret int) {
		if i == len(s) {
			return 0
		}
		if cache[i] != 0 {
			return cache[i]
		}
		defer func() {
			cache[i] = ret
		}()
		if s[i] == '0' {
			return -1
		}
		curr := 0
		ret = -1
		for j := i; j < len(s); j++ {
			curr = curr*2 + int(s[j]-'0')
			if powerOfFive(curr) { // try
				tmp := helper(j + 1)
				if tmp != -1 {
					if ret == -1 || tmp+1 < ret {
						ret = tmp + 1
					}
				}
			}
		}
		return ret
	}
	return helper(0)
}
