package main

func minimumOperations4(num string) int {
	helper := func(a, b int) int {
		var ret int
		var found1 bool
		for i := len(num) - 1; i >= 0; i-- {
			if !found1 {
				if num[i] == byte(a+'0') {
					found1 = true
				} else {
					ret += 1
				}
				continue
			}
			if num[i] == byte(b+'0') {
				return ret
			}
			ret += 1
		}
		return len(num)
	}
	ret := min(min(helper(0, 5), helper(0, 0)), min(helper(5, 2), helper(5, 7)))
	var tmp int
	for _, c := range num {
		if c != '0' {
			tmp += 1
		}
	}
	return min(ret, tmp)
}
