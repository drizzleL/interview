package main

func doesAliceWin(s string) bool {
	var cnt int
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			cnt += 1
		}
	}
	return cnt != 0
}
