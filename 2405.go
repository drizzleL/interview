package main

func partitionString(s string) int {
	if len(s) == 0 {
		return 0
	}
	ret := 1
	dict := [26]bool{}
	for i := 0; i < len(s); i++ {
		num := int(s[i] - 'a')
		if dict[num] {
			ret += 1
			dict = [26]bool{}
		}
		dict[num] = true
	}
	return ret
}
