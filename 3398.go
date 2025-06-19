package main

func minLength2(s string, numOps int) int {
	var cnts []int
	cnt := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt += 1
			continue
		}
		cnts = append(cnts, cnt)
		cnt = 1
	}
	cnts = append(cnts, cnt)
	checkSingle := func(x byte) bool {
		var ret int
		for i := 0; i < len(s); i++ {
			switch s[i] - '0' {
			case x:
			default:
				ret += 1
			}
			x = 1 - x
		}
		return ret <= numOps
	}
	check := func(x int) bool {
		if x == 1 {
			return checkSingle(0) || checkSingle(1)
		}
		ops := numOps
		for _, cnt := range cnts {
			ops -= cnt / (x + 1)
		}
		return ops >= 0
	}
	l, r := 1, len(s)
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
