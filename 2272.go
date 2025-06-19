package main

func largestVariance(s string) int {
	dict := make([]bool, 26)
	var cc []byte
	for _, c := range s {
		if !dict[c-'a'] {
			cc = append(cc, byte(c))
		}
		dict[c-'a'] = true
	}
	helper := func(a, b byte) int {
		var ret, cnt int
		var flag, leftFlag bool
		for i := 0; i < len(s); i++ {
			var val int
			switch s[i] {
			case a:
				val = 1
			case b:
				val = -1
			}
			if val >= 0 {
				cnt += val
				if flag {
					ret = max(ret, cnt)
				}
				continue
			}
			flag = true
			cnt += val
			if leftFlag {
				cnt += 1
				leftFlag = false
			}
			if cnt >= 0 {
				ret = max(ret, cnt)
				continue
			}
			cnt = -1
			leftFlag = true
		}
		return ret
	}
	var ret int
	for i := 0; i < len(cc); i++ {
		for j := 0; j < len(cc); j++ {
			if i == j {
				continue
			}
			ret = max(ret, helper(cc[i], cc[j]))
		}
	}
	return ret
}
