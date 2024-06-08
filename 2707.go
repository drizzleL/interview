package main

func minExtraChar(s string, dictionary []string) int {
	type trieNode struct {
		flag     bool
		children [26]*trieNode
	}
	root := &trieNode{}
	for _, w := range dictionary {
		n := root
		for _, c := range w {
			if n.children[c-'a'] == nil {
				n.children[c-'a'] = &trieNode{}
			}
			n = n.children[c-'a']
		}
		n.flag = true
	}
	cache := make([]int, len(s))
	for i := range cache {
		cache[i] = -1
	}
	find := func(i int) []int {
		var ret []int
		node := root
		for step := 0; node != nil && i+step <= len(s); step++ {
			if node.flag {
				ret = append(ret, i+step)
			}
			if i+step < len(s) {
				node = node.children[s[i+step]-'a']
			}
		}
		return ret
	}
	var helper func(i int) int
	helper = func(i int) (ret int) {
		if i >= len(s) {
			return 0
		}
		if cache[i] != -1 {
			return cache[i]
		}
		defer func() {
			cache[i] = ret
		}()
		ret = 1 + helper(i+1)
		for _, next := range find(i) {
			ret = min(ret, helper(next))
		}
		return
	}
	return helper(0)
}
