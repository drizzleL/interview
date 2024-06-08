package main

func diStringMatch(s string) []int {
	l, r := 0, len(s)
	var ret []int
	for _, c := range s {
		switch c {
		case 'I':
			ret = append(ret, l)
			l++
		case 'D':
			ret = append(ret, r)
			r--
		}
	}
	ret = append(ret, l)
	return ret
}
