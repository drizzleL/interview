package main

func longestValidSubstring(word string, forbidden []string) int {
	type trieNode struct {
		children [26]*trieNode
		flag     bool
	}
	root := &trieNode{}
	for _, forb := range forbidden {
		cur := root
		for i := len(forb) - 1; i >= 0; i-- {
			c := forb[i]
			if cur.children[c-'a'] == nil {
				cur.children[c-'a'] = &trieNode{}
			}
			cur = cur.children[c-'a']
		}
		cur.flag = true
	}
	check := func(i, j int) bool {
		if i > j {
			return true
		}
		cur := root
		for k := j; k >= i; k-- {
			c := word[k]
			if cur.children[c-'a'] == nil {
				return true
			}
			cur = cur.children[c-'a']
			if cur.flag {
				return false
			}
		}
		return true
	}
	var ret int
	for i, j := 0, 0; j < len(word); j++ {
		for !check(i, j) {
			i++
		}
		ret = max(ret, j-i+1)
	}
	return ret
}
