package main

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	ret := make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, n)
		for j := range ret[i] {
			ret[i][j] = -1
		}
	}
	up, down, left, right := 0, m-1, 0, n-1
	var i, j int
	j = -1
	for head != nil {
		for head != nil {
			j += 1
			ret[i][j] = head.Val
			head = head.Next
			if j == right {
				break
			}
		}
		up += 1
		for head != nil {
			i += 1
			ret[i][j] = head.Val
			head = head.Next
			if i == down {
				break
			}
		}
		right -= 1
		for head != nil {
			j -= 1
			ret[i][j] = head.Val
			head = head.Next
			if j == left {
				break
			}
		}
		down -= 1
		for head != nil {
			i -= 1
			ret[i][j] = head.Val
			head = head.Next
			if i == up {
				break
			}
		}
		left += 1
	}
	return ret
}
