package main

func validStrings(n int) []string {
	var ret []string
	b := make([]byte, n)
	var helper func(i int)
	helper = func(i int) {
		if i == n {
			ret = append(ret, string(b))
			return
		}
		b[i] = '1'
		helper(i + 1)
		if i != 0 && b[i-1] == '0' {
			return
		}
		b[i] = '0'
		helper(i + 1)
	}
	helper(0)
	return ret
}
