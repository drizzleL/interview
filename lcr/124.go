package main

func deduceTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var idx int
	for i := range inorder {
		if inorder[i] == preorder[0] {
			idx = i
		}
	}
	root.Left = deduceTree(preorder[1:idx+1], inorder[:idx])
	root.Right = deduceTree(preorder[idx+1:], inorder[idx+1:])
	return root
}
