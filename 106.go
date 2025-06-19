package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	node := &TreeNode{
		Val: val,
	}
	var idx int
	for ; inorder[idx] != val; idx++ {
	}
	node.Left = buildTree(inorder[:idx], postorder[:idx])
	node.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return node
}
