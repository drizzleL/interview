package main

func verifyTreeOrder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	rootVal := postorder[len(postorder)-1]
	var mid int
	for i := 0; i < len(postorder)-1; i++ {
		if postorder[i] > rootVal {
			mid = i
			break
		}
	}
	for i := mid + 1; i < len(postorder)-1; i++ {
		if postorder[i] < rootVal {
			return false
		}
	}
	return verifyTreeOrder(postorder[:mid]) && verifyTreeOrder(postorder[mid:len(postorder)-1])
}
