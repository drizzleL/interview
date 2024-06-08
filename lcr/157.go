package main

import "sort"

func goodsOrder(goods string) []string {
	b := []byte(goods)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	vis := make([]bool, len(b))
	var ret []string
	var helper func(i int, tmp []byte)
	helper = func(i int, tmp []byte) {
		if i == len(b) {
			ret = append(ret, string(tmp))
			return
		}
		for j := range b {
			if vis[j] {
				continue
			}
			if j != 0 && b[j] == b[j-1] && !vis[j-1] {
				continue
			}
			vis[j] = true
			helper(i+1, append(tmp, b[j]))
			vis[j] = false
		}
	}
	helper(0, nil)
	return ret
}
