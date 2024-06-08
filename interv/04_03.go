package main

func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return nil
	}
	nodes := []*TreeNode{tree}
	var ret []*ListNode
	for len(nodes) != 0 {
		var next []*TreeNode
		dummyHead := &ListNode{}
		tmp := dummyHead
		for _, n := range nodes {
			tmp.Next = &ListNode{Val: n.Val}
			tmp = tmp.Next
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		ret = append(ret, dummyHead.Next)
		nodes = next
	}
	return ret
}
