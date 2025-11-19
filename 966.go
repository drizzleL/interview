package main

func spellchecker(wordlist []string, queries []string) []string {
	perfect := map[string]bool{}
	caseIns := map[string]string{}
	vowelMis := map[string]string{}
	lower := func(b byte) byte {
		if b >= 'A' && b <= 'Z' {
			return b + ('a' - 'A')
		}
		return b
	}
	helper1 := func(s string) string {
		var b []byte
		for _, c := range s {
			b = append(b, lower(byte(c)))
		}
		return string(b)
	}
	helper2 := func(s string) string {
		var b []byte
		for _, c := range s {
			switch c {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				b = append(b, '*')
			default:
				b = append(b, lower(byte(c)))
			}
		}
		return string(b)
	}
	for _, w := range wordlist {
		perfect[w] = true
		w1 := helper1(w)
		if _, ok := caseIns[w1]; !ok {
			caseIns[w1] = w
		}
		w2 := helper2(w)
		if _, ok := vowelMis[w2]; !ok {
			vowelMis[w2] = w
		}
	}
	var ret []string
	for _, q := range queries {
		if perfect[q] {
			ret = append(ret, q)
			continue
		}
		if v, ok := caseIns[helper1(q)]; ok {
			ret = append(ret, v)
			continue
		}
		if v, ok := vowelMis[helper2(q)]; ok {
			ret = append(ret, v)
			continue
		}
		ret = append(ret, "")
	}
	return ret
}
