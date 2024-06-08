package main

func insert(aNode *Node, x int) *Node {
	n := &Node{
		Val: x,
	}
	if aNode == nil {
		n.Next = n
		return n
	}
	curr := aNode
	for curr.Next != aNode {
		next := curr.Next
		if curr.Val <= x && x <= next.Val {
			break
		}
		if curr.Val > next.Val {
			if x > curr.Val || x < next.Val { // found in mid
				break
			}
		}
		curr = curr.Next
	}
	n.Next = curr.Next
	curr.Next = n
	return aNode
}
