package main

import "strings"

type FolderNode struct {
	Children map[string]*FolderNode
	Flag     bool
}

func removeSubfolders2(folder []string) []string {
	type node struct {
		children map[string]*node
		isEnd    bool
	}
	root := &node{
		children: map[string]*node{},
	}
	for _, fol := range folder {
		path := strings.Split(fol, "/")
		n := root
		for _, name := range path {
			if n.children[name] == nil {
				n.children[name] = &node{
					children: map[string]*node{},
				}
			}
			n = n.children[name]
		}
		n.isEnd = true
	}
	var ret []string
	var dfs func(n *node, path []string)
	dfs = func(n *node, path []string) {
		if n.isEnd {
			ret = append(ret, strings.Join(path, "/"))
			return
		}
		for name, child := range n.children {
			dfs(child, append(path, name))
		}
	}
	dfs(root, nil)
	return ret
}
func removeSubfolders(folder []string) []string {
	var ret []string
	root := &FolderNode{
		Children: map[string]*FolderNode{},
	}
	for _, fld := range folder {
		node := root
		for i, name := range strings.Split(fld, "/") {
			if i == 0 {
				continue
			}
			if node.Children[name] == nil {
				node.Children[name] = &FolderNode{
					Children: map[string]*FolderNode{},
				}
			}
			node = node.Children[name]
		}
		node.Flag = true
	}
	var dfs func(node *FolderNode, path []string)
	dfs = func(node *FolderNode, path []string) {
		if node.Flag {
			ret = append(ret, strings.Join(path, "/"))
			return
		}
		for name, child := range node.Children {
			dfs(child, append(path, name))
		}
	}
	dfs(root, []string{""})
	return ret
}
