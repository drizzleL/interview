package main

func minFlips2(a int, b int, c int) int {
	var ret int
	for ; a != 0 || b != 0 || c != 0; a, b, c = a>>1, b>>1, c>>1 {
		if c&1 == 0 {
			ret += a & 1
			ret += b & 1
			continue
		}
		ret += min(1, 1-((a&1)|(b&1)))
	}
	return ret
}
