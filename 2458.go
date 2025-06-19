package main

func treeQueries(root *TreeNode, queries []int) []int {
	nodeMaxHeight := map[int]int{} // node to height
	nodeLvl := map[int]int{}
	lvlDict := map[int][]int{} // level to nodes
	var dfs func(node *TreeNode, lvl int) int
	dfs = func(node *TreeNode, lvl int) int {
		if node == nil {
			return 0
		}
		nodeLvl[node.Val] = lvl
		tmp := max(dfs(node.Left, lvl+1), dfs(node.Right, lvl+1))
		height := lvl + tmp
		nodeMaxHeight[node.Val] = height
		switch len(lvlDict[lvl]) {
		case 0:
			lvlDict[lvl] = append(lvlDict[lvl], node.Val)
		case 1:
			if nodeMaxHeight[lvlDict[lvl][0]] > height {
				lvlDict[lvl] = append(lvlDict[lvl], node.Val)
			} else {
				lvlDict[lvl] = append([]int{node.Val}, lvlDict[lvl]...)
			}
		case 2:
			if height >= nodeMaxHeight[lvlDict[lvl][0]] {
				lvlDict[lvl][1] = lvlDict[lvl][0]
				lvlDict[lvl][0] = node.Val
			} else if height > nodeMaxHeight[lvlDict[lvl][1]] {
				lvlDict[lvl][1] = node.Val
			}
		}
		return tmp + 1
	}
	dfs(root, 0)
	ret := make([]int, len(queries))
	for i, q := range queries {
		lvl := nodeLvl[q]
		if len(lvlDict[lvl]) == 1 { // this q it self in this level
			ret[i] = lvl - 1
			continue
		}
		if lvlDict[lvl][0] == q {
			ret[i] = nodeMaxHeight[lvlDict[lvl][1]]
		} else {
			ret[i] = nodeMaxHeight[lvlDict[lvl][0]]
		}
	}
	return ret
}
