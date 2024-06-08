package main

import (
	"sort"
	"strings"
)

func accountsMerge(accounts [][]string) [][]string {
	type Item struct {
		Name string
		Mail string
	}
	var parent []int
	var find func(i int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}
	var items []*Item
	union := func(a, b int) {
		fa, fb := find(a), find(b)
		if fa > fb {
			fa, fb = fb, fa
		}
		parent[fb] = fa
	}
	dict := map[string]int{}
	for _, account := range accounts {
		for i := 1; i < len(account); i++ {
			if _, ok := dict[account[i]]; !ok {
				item := &Item{
					Name: account[0],
					Mail: account[i],
				}
				parent = append(parent, len(parent))
				dict[account[i]] = len(items)
				items = append(items, item)
			}
			if i != 1 {
				union(dict[account[1]], dict[account[i]])
			}
		}
	}
	groups := map[int][]*Item{}
	for i, item := range items {
		groups[find(i)] = append(groups[find(i)], item)
	}
	var ret [][]string
	for _, g := range groups {
		var tmp []string
		for i := 0; i < len(g); i++ {
			tmp = append(tmp, g[i].Mail)
		}
		sort.Strings(tmp)
		tmp = append([]string{g[0].Name}, tmp...)
		ret = append(ret, tmp)
	}
	sort.Slice(ret, func(i, j int) bool {
		return strings.Compare(ret[i][0], ret[j][0]) < 0
	})
	return ret
}
