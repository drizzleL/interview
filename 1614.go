package main

func maxDepth(s string) int {
	var ret int
	var tmp int
	var cnt int
	for _, c := range s {
		switch c {
		case '(':
			tmp += 1
			cnt = max(cnt, tmp)
		case ')':
			tmp -= 1
			if tmp == 0 {
				ret = max(ret, cnt)
				cnt = 0
			}
		}
	}
	return ret
}
