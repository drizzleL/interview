package main

func minRemoveToMakeValid(s string) string {
	check := make([]bool, len(s))
	lefts := []int{}
	for i, c := range s {
		switch c {
		case '(':
			lefts = append(lefts, i)
		case ')':
			if len(lefts) != 0 {
				check[i] = true
				top := lefts[len(lefts)-1]
				lefts = lefts[:len(lefts)-1]
				check[top] = true
			}
		default:
			check[i] = true
		}
	}
	var ret []byte
	for i, c := range check {
		if c {
			ret = append(ret, s[i])
		}
	}
	return string(ret)
}
