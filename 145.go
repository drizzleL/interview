package main

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := postorderTraversal(root.Left)
	ret = append(ret, postorderTraversal(root.Right)...)
	ret = append(ret, root.Val)
	return ret
}
