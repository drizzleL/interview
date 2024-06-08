package main

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret int
	nodes := []*TreeNode{root}
	nums := []int{0}
	for len(nodes) != 0 {
		var nextnodes []*TreeNode
		var nextnums []int
		for i, n := range nodes {
			if n.Left == nil && n.Right == nil {
				ret += nums[i] + n.Val
				continue
			}
			if n.Left != nil {
				nextnums = append(nextnums, (nums[i]+n.Val)*10)
				nextnodes = append(nextnodes, n.Left)
			}
			if n.Right != nil {
				nextnums = append(nextnums, (nums[i]+n.Val)*10)
				nextnodes = append(nextnodes, n.Right)
			}
		}
		nodes = nextnodes
		nums = nextnums
	}
	return ret
}
