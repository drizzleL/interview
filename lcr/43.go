package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root      *TreeNode
	lastnodes []*TreeNode
	nextnodes []*TreeNode
}

func CBTConstructor(root *TreeNode) CBTInserter {
	nodes := []*TreeNode{root}
	lastnodes := []*TreeNode{root}
	var nextnodes []*TreeNode
	for len(nodes) != 0 {
		var next []*TreeNode
		for _, n := range nodes {
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		if len(next) != 0 || len(nodes) == len(lastnodes)*2 {
			lastnodes = nodes
			nextnodes = next
		}
		nodes = next
	}
	return CBTInserter{
		root:      root,
		lastnodes: lastnodes,
		nextnodes: nextnodes,
	}
}

func (this *CBTInserter) Insert(v int) int {
	lastnode := this.lastnodes[len(this.nextnodes)/2]
	n := &TreeNode{Val: v}
	if lastnode.Left == nil {
		lastnode.Left = n
	} else {
		lastnode.Right = n
	}
	this.nextnodes = append(this.nextnodes, n)
	if len(this.nextnodes) == len(this.lastnodes)*2 { // next full
		this.lastnodes, this.nextnodes = this.nextnodes, nil
	}
	return lastnode.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
