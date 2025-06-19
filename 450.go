package main

func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		maxVal := root.Left.Val
		for node := root.Left; node != nil; node = node.Right {
			maxVal = max(maxVal, node.Val)
		}
		root.Left = deleteNode2(root.Left, maxVal)
		root.Val = maxVal
	} else if root.Val > key {
		root.Left = deleteNode2(root.Left, key)
	} else {
		root.Right = deleteNode2(root.Right, key)
	}
	return root
}
