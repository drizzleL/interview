package main

func getDirections(root *TreeNode, startValue int, destValue int) string {
	var find func(node *TreeNode, val int) ([]byte, bool)
	find = func(node *TreeNode, val int) ([]byte, bool) {
		if node == nil {
			return nil, false
		}
		if node.Val == val {
			return nil, true
		}
		lpath, found := find(node.Left, val)
		if found {
			return append(lpath, 'L'), true
		}
		rpath, found := find(node.Right, val)
		if found {
			return append(rpath, 'R'), true
		}
		return nil, false
	}
	startPath, _ := find(root, startValue)
	destPath, _ := find(root, destValue)
	reverse := func(b []byte) {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}
	reverse(startPath)
	reverse(destPath)
	for len(startPath) > 0 && len(destPath) > 0 && startPath[0] == destPath[0] {
		startPath = startPath[1:]
		destPath = destPath[1:]
	}
	var ret []byte
	for i := 0; i < len(startPath); i++ {
		ret = append(ret, 'U')
	}
	ret = append(ret, destPath...)
	return string(ret)
}
