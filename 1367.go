package main

func isSubPath(head *ListNode, root *TreeNode) bool {
	var dfs func(head *ListNode, root *TreeNode) bool
	dfs = func(head *ListNode, root *TreeNode) bool {
		if head == nil {
			return true
		}
		if root == nil {
			return false
		}
		return head.Val == root.Val && (dfs(head.Next, root.Left) || dfs(head.Next, root.Right))
	}
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}
