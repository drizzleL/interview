package main

func findWords(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])
	type trieNode struct {
		children [26]*trieNode
		flag     int
	}
	root := &trieNode{
		flag: -1,
	}
	for i, w := range words {
		node := root
		for j := 0; j < len(w); j++ {
			idx := int(w[j] - 'a')
			if node.children[idx] == nil {
				node.children[idx] = &trieNode{flag: -1}
			}
			node = node.children[idx]
		}
		node.flag = i
	}
	var ret []string
	var dfs func(i, j int, node *trieNode)
	dfs = func(i, j int, node *trieNode) {
		if board[i][j] == '#' {
			return
		}
		c := board[i][j]
		idx := int(c - 'a')
		if node.children[idx] == nil {
			return
		}
		board[i][j] = '#'
		node = node.children[idx]
		if node.flag != -1 {
			ret = append(ret, words[node.flag])
			node.flag = -1
		}
		for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			i2, j2 := i+dir[0], j+dir[1]
			if i2 < 0 || j2 < 0 || i2 >= m || j2 >= n {
				continue
			}
			dfs(i2, j2, node)
		}
		board[i][j] = c
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, root)
		}
	}
	return ret
}
