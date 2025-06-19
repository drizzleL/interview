package main

func allPossibleFBT(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{{}}
	}
	var ret []*TreeNode
	for left := 1; left < n-1; left += 2 {
		right := n - 1 - left
		ll := allPossibleFBT(left)
		rr := allPossibleFBT(right)
		for lidx := 0; lidx < len(ll); lidx++ {
			for ridx := 0; ridx < len(rr); ridx++ {
				root := &TreeNode{}
				root.Left = ll[lidx]
				root.Right = rr[ridx]
				ret = append(ret, root)
			}
		}
	}
	return ret
}
