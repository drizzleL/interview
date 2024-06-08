package main

func minimumDeletions(s string) int {
	var aCnt, bCnt int
	for _, c := range s {
		if c == 'a' {
			aCnt += 1
		}
	}
	ret := aCnt
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			aCnt -= 1
		case 'b':
			bCnt += 1
		}
		ret = min(aCnt+bCnt, ret)
	}
	return ret
}
