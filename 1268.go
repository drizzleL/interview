package main

import "sort"

func suggestedProducts(products []string, searchWord string) [][]string {
	type trieNode struct {
		tops     []string
		children [26]*trieNode
	}
	root := &trieNode{}
	sort.Strings(products)
	for _, p := range products {
		node := root
		for _, c := range p {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &trieNode{}
			}
			node = node.children[c-'a']
			node.tops = append(node.tops, p)
			if len(node.tops) > 3 {
				node.tops = node.tops[:3]
			}
		}
	}
	ret := make([][]string, len(searchWord))
	for i, c := range searchWord {
		root = root.children[c-'a']
		if root == nil {
			break
		}
		ret[i] = root.tops
	}
	return ret
}
