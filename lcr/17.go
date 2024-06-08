package main

func minWindow(s string, t string) string {
	dict := map[byte]int{}
	for _, c := range t {
		dict[byte(c)]++
	}
	var ret string
	candidate := len(dict)
	for i, j := 0, 0; j < len(s); {
		for ; candidate != 0 && j < len(s); j++ {
			if _, ok := dict[s[j]]; !ok { //nonsense
				continue
			}
			dict[s[j]]--
			if dict[s[j]] == 0 {
				candidate--
			}
		}
		for ; i < j && candidate == 0; i++ {
			if ret == "" || j-i < len(ret) {
				ret = s[i:j]
			}
			if _, ok := dict[s[i]]; !ok {
				continue
			}
			dict[s[i]]++
			if dict[s[i]] > 0 {
				candidate++
				i++
				break
			}
		}
	}
	return ret
}
