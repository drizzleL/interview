package main

func createBinaryTree(descriptions [][]int) *TreeNode {
	if len(descriptions) == 0 {
		return nil
	}
	dict := map[int][][]int{}
	dictIsChild := map[int]bool{}
	for _, desc := range descriptions {
		dict[desc[0]] = append(dict[desc[0]], desc)
		dictIsChild[desc[1]] = true
	}
	var parentIdx int
	for k := range dict {
		if !dictIsChild[k] {
			parentIdx = k
			break
		}
	}
	var helper func(parent int) *TreeNode
	helper = func(parent int) *TreeNode {
		ret := &TreeNode{
			Val: parent,
		}
		for _, child := range dict[parent] {
			if child[2] == 1 {
				ret.Left = helper(child[1])
			} else {
				ret.Right = helper(child[1])
			}
		}
		return ret
	}
	return helper(parentIdx)
}
