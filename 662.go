package main

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	type node struct {
		n  *TreeNode
		id int
	}
	nodes := []*node{{n: root, id: 1}}
	var ret int
	for lvl := 1; len(nodes) != 0; lvl++ {
		var next []*node
		ret = max(ret, nodes[len(nodes)-1].id-nodes[0].id+1)
		for _, n := range nodes {
			if n.n.Left != nil {
				next = append(next, &node{n: n.n.Left, id: n.id*2 - 1})
			}
			if n.n.Right != nil {
				next = append(next, &node{n: n.n.Right, id: n.id * 2})
			}
		}
		nodes = next
	}
	return ret
}
