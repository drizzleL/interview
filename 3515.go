package main

// func treeQueries2(n int, edges [][]int, queries [][]int) []int {
// 	var ret []int
// 	edDict := make([][][2]int, n+1)
// 	for _, ed := range edges {
// 		edDict[ed[0]] = append(edDict[ed[0]], [2]int{ed[1], ed[2]})
// 		edDict[ed[1]] = append(edDict[ed[1]], [2]int{ed[0], ed[2]})
// 	}
// 	seen := make([]bool, n+1)
// 	seen[1] = true
// 	var dfs func(x int) [][2]int
// 	dfs = func(x int) [][2]int {
// 		var ret [][2]int
// 		for _, next := range edDict[x] {
// 			nextIdx := next[0]
// 			if seen[nextIdx] {
// 				continue
// 			}
// 			seen[nextIdx] = true
// 		}
// 		return ret
// 	}
// 	paths := dfs(1)
// 	var bits []*biTree
// 	biTreeDict := make([]int, )
// 	for _, path := range paths {

// 	}
// 	for _, q := range queries {
// 		if q[0] == 1 {

//			} else {
//				if q[1] == 1 {
//					ret = append(ret, 0)
//					continue
//				}
//				ret = append(ret, 0)
//			}
//		}
//		return nil
//	}
func NewBiTree(size int) biTree {
	return biTree(make([]int, size+1))
}

type biTree []int

func (bit biTree) Query(idx int) int {
	idx += 1
	var ret int
	for idx > 0 {
		ret += bit[idx]
		idx -= lowbit(idx)
	}
	return ret
}

func (bit biTree) Update(idx int, val int) {
	idx += 1
	for idx < len(bit) {
		bit[idx] += val
		idx += lowbit(idx)
	}
}

func lowbit(x int) int {
	return x & (-x)
}
