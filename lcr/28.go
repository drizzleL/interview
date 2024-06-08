package main

type Node struct {
	Val    int
	Prev   *Node
	Next   *Node
	Child  *Node
	Random *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	pre := &Node{}
	node := pre
	nodes := []*Node{root}
	for len(nodes) != 0 {
		top := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		node.Next, top.Prev = top, node
		node = top
		if node.Next != nil {
			nodes = append(nodes, node.Next)
			node.Next = nil
		}
		if node.Child != nil {
			nodes = append(nodes, node.Child)
			node.Child = nil
		}
	}
	pre.Next.Prev = nil
	return pre.Next
}
