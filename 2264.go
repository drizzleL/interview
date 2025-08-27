package main

func largestGoodInteger(num string) string {
	var cnt int
	var lastc rune
	var ret rune
	check := func() {
		if lastc <= ret {
			return
		}
		if cnt < 3 {
			return
		}
		ret = lastc
	}
	for _, c := range num {
		if c == lastc {
			cnt += 1
			continue
		}
		check()
		lastc = c
		cnt = 1
	}
	check()
	if ret != 0 {
		return string([]rune{ret, ret, ret})
	}
	return ""
}
