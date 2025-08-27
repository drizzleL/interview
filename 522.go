package main

func findLUSlength(strs []string) int {
	type trieNode struct {
		cnt      int
		children [26]*trieNode
	}
	root := &trieNode{}
	var helper func(node *trieNode, str string, i int)
	helper = func(node *trieNode, str string, i int) {
		if i == len(str) {
			return
		}
		helper(node, str, i+1)
		if node.children[str[i]-'a'] == nil {
			node.children[str[i]-'a'] = &trieNode{}
		}
		node.children[str[i]-'a'].cnt++
		helper(node.children[str[i]-'a'], str, i+1)
	}
	for _, str := range strs {
		helper(root, str, 0)
	}
	ret := -1
	var dfs func(node *trieNode, v int)
	dfs = func(node *trieNode, v int) {
		if node.cnt == 1 {
			ret = max(ret, v)
		}
		for _, child := range node.children {
			if child == nil {
				continue
			}
			dfs(child, v+1)
		}
	}
	dfs(root, 0)
	return ret
}
