package main

func generateParenthesis(n int) []string {
	var ret []string
	var helper func(l, r int, tmp []byte)
	helper = func(l, r int, tmp []byte) {
		if l == n && r == n {
			ret = append(ret, string(tmp))
			return
		}
		if l < n {
			helper(l+1, r, append(tmp, '('))
		}
		if r < l {
			helper(l, r+1, append(tmp, ')'))
		}
	}
	helper(0, 0, nil)
	return ret
}
