package main

func minSwaps2(s string) int {
	var cnt int
	var ret int
	for _, c := range s {
		switch c {
		case '[':
			cnt += 1
		case ']':
			if cnt == 0 {
				ret += 1
				cnt += 1
			} else {
				cnt -= 1
			}
		}
	}
	return ret
}
