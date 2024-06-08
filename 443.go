package main

import "strconv"

func compress(chars []byte) int {
	var last byte
	var cnt int
	var j int
	helper := func() {
		if cnt == 0 {
			return
		}
		chars[j] = last
		j += 1
		if cnt == 1 {
			return
		}
		s := strconv.Itoa(cnt)
		for i := 0; i < len(s); i++ {
			chars[j] = s[i]
			j += 1
		}
	}
	for i := 0; i < len(chars); i++ {
		if chars[i] != last {
			helper()
			last = chars[i]
			cnt = 1
			continue
		}
		cnt += 1
	}
	helper()
	return j
}
