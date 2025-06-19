package main

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	seen := make([]bool, n)
	for _, c := range leftChild {
		if c == -1 {
			continue
		}
		seen[c] = true
	}
	for _, c := range rightChild {
		if c == -1 {
			continue
		}
		if seen[c] { // contains same child
			return false
		}
		seen[c] = true
	}
	root := -1
	for k, v := range seen {
		if v {
			continue
		}
		if root != -1 { // multiple roots
			return false
		}
		root = k
	}
	seen = make([]bool, n)
	var dfs func(node int)
	dfs = func(node int) {
		if node == -1 {
			return
		}
		seen[node] = true
		dfs(leftChild[node])
		dfs(rightChild[node])
	}
	dfs(root)
	for _, v := range seen {
		if !v {
			return false
		}
	}
	return true
}
