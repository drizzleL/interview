package main

func minFlipsMonoIncr(s string) int {
	if len(s) == 0 {
		return 0
	}
	ones := 0
	for _, c := range s {
		if c == '1' {
			ones++
		}
	}
	ret := min(len(s)-ones, ones) // all zero or all one
	rightzeros := len(s) - ones
	var leftones int
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			leftones++
		} else {
			rightzeros--
		}
		if (leftones + rightzeros) < ret {
			ret = leftones + rightzeros
		}
	}
	return ret
}
