package main

func maxDistance3(s string, k int) int {
	var extra int
	var s1, s2 int
	helper := func(s1, s2 int, extra int) int {
		ret := abs(s1) + abs(s2)
		change := min(k, extra)
		ret += change * 2
		ret += k
		return ret
	}
	var ret int
	for _, c := range s {
		switch c {
		case 'N':
			if s1 < 0 {
				extra += 1
			}
			s1 += 1
		case 'S':
			if s1 > 0 {
				extra += 1
			}
			s1 -= 1
		case 'E':
			if s2 < 0 {
				extra += 1
			}
			s2 += 1
		case 'W':
			if s2 > 0 {
				extra += 1
			}
			s2 -= 1
		}
		ret = max(ret, helper(s1, s2, extra))
	}
	return ret
}
