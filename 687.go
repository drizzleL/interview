package main

func longestUnivaluePath(root *TreeNode) int {
	var ret int
	var helper func(n *TreeNode) int
	helper = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		cnt := 1
		bothCnt := 1
		lcnt := helper(n.Left)
		rcnt := helper(n.Right)
		var tail int
		if n.Left != nil && n.Left.Val == n.Val {
			bothCnt += lcnt
			if lcnt > tail {
				tail = lcnt
			}
		}
		if n.Right != nil && n.Right.Val == n.Val {
			bothCnt += rcnt
			if rcnt > tail {
				tail = rcnt
			}
		}
		if bothCnt > ret {
			ret = bothCnt
		}
		return cnt + tail
	}
	helper(root)
	if ret == 0 {
		return 0
	}
	return ret - 1
}
