package main

func BSTSequences(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{nil}
	}
	var ret [][]int
	for _, v := range combine(BSTSequences(root.Left), BSTSequences(root.Right)) {
		ret = append(ret, append([]int{root.Val}, v...))
	}
	return ret
}

func mergeArr(a, b []int) [][]int {
	var ret [][]int
	var helper func(i, j int, tmp []int)
	helper = func(i, j int, tmp []int) {
		if i == len(a) && j == len(b) {
			dst := make([]int, len(tmp))
			copy(dst, tmp)
			ret = append(ret, dst)
			return
		}
		if i != len(a) {
			helper(i+1, j, append(tmp, a[i]))
		}
		if j != len(b) {
			helper(i, j+1, append(tmp, b[j]))
		}
	}
	helper(0, 0, nil)
	return ret
}

func combine(a, b [][]int) [][]int {
	if len(a) == 0 && len(b) == 0 {
		return [][]int{nil}
	}
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	var ret [][]int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			ret = append(ret, mergeArr(a[i], b[j])...)
		}
	}
	return ret
}
