package main

import (
	"sort"
	"strings"
)

func deleteDuplicateFolder(paths [][]string) [][]string {
	type trieNode struct {
		serial   string
		children map[string]*trieNode
	}
	root := &trieNode{
		children: map[string]*trieNode{},
	}
	for _, path := range paths {
		node := root
		for _, p := range path {
			if _, ok := node.children[p]; !ok {
				node.children[p] = &trieNode{
					children: map[string]*trieNode{},
				}
			}
			node = node.children[p]
		}
	}
	freq := map[string]int{}
	var dfs func(node *trieNode)
	dfs = func(node *trieNode) {
		if len(node.children) == 0 {
			return
		}
		var childNames []string
		for cName, child := range node.children {
			dfs(child)
			childNames = append(childNames, cName+"("+child.serial+")")
		}
		sort.Strings(childNames)
		node.serial = strings.Join(childNames, ",")
		freq[node.serial] += 1
	}
	dfs(root)
	var ret [][]string
	var helper func(node *trieNode, curr []string)
	helper = func(node *trieNode, curr []string) {
		if freq[node.serial] > 1 {
			return
		}
		if len(curr) > 0 {
			tmp := make([]string, len(curr))
			copy(tmp, curr)
			ret = append(ret, tmp)
		}
		for name, child := range node.children {
			helper(child, append(curr, name))
		}
	}
	helper(root, nil)
	return ret
}
