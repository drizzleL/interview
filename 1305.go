package main

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	g1, g2 := NewGetter(root1), NewGetter(root2)
	var ret []int
	for g1.Top() != nil && g2.Top() != nil {
		top1, top2 := g1.Top(), g2.Top()
		if top1.Val < top2.Val {
			ret = append(ret, g1.Pop().Val)
		} else {
			ret = append(ret, g2.Pop().Val)
		}
	}
	for g1.Top() != nil {
		ret = append(ret, g1.Pop().Val)
	}
	for g2.Top() != nil {
		ret = append(ret, g2.Pop().Val)
	}
	return ret
}

type Getter struct {
	q    []*TreeNode
	used []bool
}

func (g *Getter) Pop() *TreeNode {
	top := g.Top()
	g.used[len(g.used)-1] = true
	n := top.Right
	for n != nil {
		g.q = append(g.q, n)
		g.used = append(g.used, false)
		n = n.Left
	}
	for len(g.used) != 0 && g.used[len(g.used)-1] {
		g.used = g.used[:len(g.used)-1]
		g.q = g.q[:len(g.q)-1]
	}
	return top
}

func (g *Getter) Top() *TreeNode {
	if len(g.q) == 0 {
		return nil
	}
	return g.q[len(g.q)-1]
}

func NewGetter(n *TreeNode) *Getter {
	g := &Getter{}
	for n != nil {
		g.q = append(g.q, n)
		g.used = append(g.used, false)
		n = n.Left
	}
	return g
}
