package main

func minFlips3(target string) int {
	var ret int
	var flag int
	for i := 0; i < len(target); i++ {
		if int(target[i]-'0') == flag {
			continue
		}
		ret += 1
		flag ^= 1
	}
	return ret
}
