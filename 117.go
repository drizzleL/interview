package main

type Node struct {
	Val      int
	Left     *Node
	Right    *Node
	Next     *Node
	Children []*Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	nodes := []*Node{root}
	for len(nodes) != 0 {
		var next []*Node
		for i, n := range nodes {
			if i != len(nodes)-1 {
				n.Next = nodes[i+1]
			}
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		nodes = next
	}
	return root
}
