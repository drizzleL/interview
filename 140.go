package main

import (
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	type trieNode struct {
		flag     bool
		children [26]*trieNode
	}
	root := &trieNode{}
	for _, w := range wordDict {
		n := root
		for _, c := range w {
			if n.children[c-'a'] == nil {
				n.children[c-'a'] = &trieNode{}
			}
			n = n.children[c-'a']
		}
		n.flag = true
	}
	find := func(i int) []int {
		var ret []int
		n := root
		for i < len(s) && n != nil {
			n = n.children[s[i]-'a']
			i++
			if n == nil {
				break
			}
			if n.flag {
				ret = append(ret, i)
			}
		}
		return ret
	}
	cache := map[int][][]string{}
	var helper func(i int) [][]string
	helper = func(i int) (ret [][]string) {
		if c, ok := cache[i]; ok {
			return c
		}
		defer func() {
			cache[i] = ret
		}()
		if i == len(s) {
			return [][]string{{}}
		}
		for _, j := range find(i) {
			pre := s[i:j]
			for _, tail := range helper(j) {
				ret = append(ret, append([]string{pre}, tail...))
			}
		}
		return ret
	}
	var ret []string
	for _, v := range helper(0) {
		ret = append(ret, strings.Join(v, " "))
	}
	return ret
}
