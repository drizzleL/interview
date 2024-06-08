package main

func balancedString(s string) int {
	dict := [128]int{}
	num := len(s) / 4
	for _, c := range s {
		dict[c] += 1
	}
	var i int
	ret := len(s)
	for j := 0; j < len(s); j++ {
		dict[s[j]]--
		for i < len(s) && dict['Q'] <= num && dict['W'] <= num && dict['E'] <= num && dict['R'] <= num {
			ret = min(ret, j-i+1)
			dict[s[i]]++
			i++
		}
	}
	return ret
}
