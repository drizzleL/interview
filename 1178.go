package main

func findNumOfValidWords(words []string, puzzles []string) []int {
	cnt := map[string]int{}
	toStr := func(w string) string {
		var dict [26]bool
		for _, c := range w {
			dict[c-'a'] = true
		}
		var b []byte
		for i, v := range dict {
			if !v {
				continue
			}
			b = append(b, byte('a'+i))
		}
		return string(b)
	}
	for _, w := range words {
		cnt[toStr(w)] += 1
	}
	var getStrs func(p string, i int, now []byte, ret *[]string)
	getStrs = func(p string, i int, now []byte, ret *[]string) {
		if i == len(p) {
			*ret = append(*ret, string(now))
			return
		}
		getStrs(p, i+1, append(now, p[i]), ret)
		getStrs(p, i+1, now, ret)
	}
	ret := make([]int, len(puzzles))
	helper := func(str string, c byte) string {
		b := []byte(str)
		for i := 0; i < len(b); i++ {
			if b[i] > c {
				b = append(b, 0)
				copy(b[i+1:], b[i:])
				b[i] = c
				return string(b)
			}
		}
		return string(append(b, c))
	}
	for i, p := range puzzles {
		var strs []string
		getStrs(toStr(p[1:]), 0, nil, &strs)
		for _, str := range strs {
			ret[i] += cnt[helper(str, p[0])]
		}
	}
	return ret
}
