package main

func permutation(S string) []string {
	vis := make([]bool, len(S))
	var ret []string
	var helper func(tmp string)
	helper = func(tmp string) {
		if len(tmp) == len(S) {
			ret = append(ret, tmp)
			return
		}
		for i := range S {
			if vis[i] {
				continue
			}
			vis[i] = true
			helper(tmp + string(S[i]))
			vis[i] = false
		}
	}
	helper("")
	return ret
}
