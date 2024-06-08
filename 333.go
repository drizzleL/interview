package main

func largestBSTSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret int
	var helper func(n *TreeNode) (int, int, int)
	helper = func(n *TreeNode) (int, int, int) {
		if n == nil {
			return 0, 0, 0
		}
		cnt1, min1, max1 := helper(n.Left)
		cnt2, min2, max2 := helper(n.Right)
		if cnt1 == -1 || cnt2 == -1 {
			return -1, 0, 0
		}
		min, max := root.Val, root.Val
		if cnt1+cnt2+1 > ret {
			ret = cnt1 + cnt2 + 1
		}
		if cnt1 != 0 {
			if min1 >= min || min2 >= min {
				return -1, 0, 0
			}
			min = min1
		}
		if cnt2 != 0 {
			if max2 <= max || max1 <= max {
				return -1, 0, 0
			}
			max = max2
		}
		return cnt1 + cnt2 + 1, min, max
	}
	helper(root)
	return ret
}
