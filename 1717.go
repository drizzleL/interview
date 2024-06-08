package main

func maximumGain(s string, x int, y int) int {
	var flag, cnta, cntb int
	if y > x {
		flag = 1
	}
	helper := func() int {
		if flag == 0 {
			return y * min(cnta, cntb)
		}
		return x * min(cnta, cntb)
	}
	var ret int
	for _, c := range s {
		if c != 'a' && c != 'b' {
			ret += helper()
			cnta, cntb = 0, 0
			continue
		}
		switch c {
		case 'a':
			if flag == 1 && cntb != 0 {
				cntb -= 1
				ret += y
				continue
			}
			cnta += 1
		case 'b':
			if flag == 0 && cnta != 0 {
				cnta -= 1
				ret += x
				continue
			}
			cntb += 1
		}
	}
	ret += helper()
	return ret
}
