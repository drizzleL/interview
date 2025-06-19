package main

import "math"

type TreeAncestor struct {
	h     int
	jumps [][]int
}

func TreeConstructor(n int, parent []int) TreeAncestor {
	h := int(math.Ceil(math.Log2(float64(n)))) + 1
	jumps := make([][]int, n)
	for i := range jumps {
		jumps[i] = make([]int, h)
		for j := range jumps[i] {
			jumps[i][j] = -1
		}
	}
	for i, p := range parent {
		jumps[i][0] = p
	}
	for j := 1; j < h; j++ {
		for i := 0; i < n; i++ {
			if jumps[i][j-1] == -1 {
				continue
			}
			jumps[i][j] = jumps[jumps[i][j-1]][j-1]
		}
	}
	return TreeAncestor{
		h:     h,
		jumps: jumps,
	}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	if node == -1 {
		return -1
	}
	if k == 0 {
		return node
	}
	if k == 1 {
		return this.jumps[node][0]
	}
	m := int(math.Floor(math.Log2(float64(k))))
	steps := int(math.Pow(2, float64(m)))
	return this.GetKthAncestor(this.jumps[node][m], k-steps)
}
