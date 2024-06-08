package main

func generateParenthesis(n int) []string {
	var ret []string
	var helper func(tmp string, l, r int)
	helper = func(tmp string, l, r int) {
		if l > n || r > l {
			return
		}
		if len(tmp) == n*2 {
			ret = append(ret, tmp)
			return
		}
		helper(tmp+"(", l+1, r)
		helper(tmp+")", l, r+1)
	}
	helper("", 0, 0)
	return ret
}
