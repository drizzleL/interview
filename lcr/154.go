package main

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	node := head
	for node != nil {
		cNode := &Node{Val: node.Val}
		node.Next, cNode.Next = cNode, node.Next
		node = node.Next.Next
	}
	node = head
	for node != nil {
		cNode := node.Next
		if node.Random != nil {
			cNode.Random = node.Random.Next
		}
		node = node.Next.Next
	}
	l1, l2 := &Node{}, &Node{}
	ret := l2
	node = head
	for node != nil {
		l1.Next = node
		l2.Next = node.Next
		l1 = l1.Next
		l2 = l2.Next
		node = l2.Next
		l1.Next = nil
		l2.Next = nil
	}

	return ret.Next
}
