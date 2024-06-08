package main

import "bytes"

func repeatLimitedString(s string, repeatLimit int) string {
	dict := [26]int{}
	var ret []byte
	for _, c := range s {
		dict[c-'a']++
	}
	cc := []int{}
	for i := 25; i >= 0; i-- {
		if dict[i] == 0 {
			continue
		}
		cc = append(cc, i)
	}
	for len(cc) != 0 {
		if dict[cc[0]] <= repeatLimit {
			tmp := bytes.Repeat([]byte{byte('a' + cc[0])}, dict[cc[0]])
			ret = append(ret, tmp...)
			cc = cc[1:]
			continue
		}
		tmp := bytes.Repeat([]byte{byte('a' + cc[0])}, repeatLimit)
		dict[cc[0]] -= repeatLimit
		ret = append(ret, tmp...)
		if len(cc) == 1 {
			break
		}
		ret = append(ret, byte('a'+cc[1]))
		dict[cc[1]] -= 1
		if dict[cc[1]] == 0 {
			cc[1] = cc[0]
			cc = cc[1:]
		}
	}
	return string(ret)
}
