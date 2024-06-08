package main

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var ret []int
	for _, child := range root.Children {
		ret = append(ret, postorder(child)...)
	}
	ret = append(ret, root.Val)
	return ret
}
