package main

func respace2(dictionary []string, sentence string) int {
	dp := make([]int, len(sentence)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] + 1
		for _, word := range dictionary {
			if i-len(word) >= 0 && word == sentence[i-len(word):i] {
				dp[i] = min(dp[i], dp[i-len(word)])
			}
		}
	}
	return dp[len(sentence)]
}
func respace(dictionary []string, sentence string) int {
	type trieNode struct {
		flag     bool
		children [26]*trieNode
	}
	root := &trieNode{}
	for _, word := range dictionary {
		node := root
		for _, c := range word {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &trieNode{}
			}
			node = node.children[c-'a']
		}
		node.flag = true
	}
	find := func(i int) (ret []int) {
		node := root
		for i < len(sentence) {
			c := sentence[i]
			node = node.children[c-'a']
			if node == nil {
				break
			}
			if node.flag {
				ret = append(ret, i+1)
			}
			i++
		}
		return ret
	}
	var helper func(i int) int
	cache := map[int]int{}
	helper = func(i int) (ret int) {
		defer func() {
			cache[i] = ret
		}()
		if c, ok := cache[i]; ok {
			return c
		}
		if i == len(sentence) {
			return 0
		}
		ret = 1 + helper(i+1)
		for _, next := range find(i) {
			ret = min(ret, helper(next))
		}
		return ret
	}
	return helper(0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
