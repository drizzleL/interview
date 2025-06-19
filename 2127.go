package main

// func maximumInvitations(favorite []int) int {
// 	closer := map[int]bool{}
// 	var closerCnt int
// 	for i, v := range favorite {
// 		if v < i {
// 			continue
// 		}
// 		if closer[i] {
// 			continue
// 		}
// 		if favorite[v] != i {
// 			continue
// 		}
// 		closer[i] = true
// 		closer[v] = true
// 		closerCnt += 1
// 	}
// 	closerExtend := closerCnt * 2
// 	ret := closerExtend
// 	helper := func(i int) {

// 	}
// 	for i, v := range favorite {
// 		helper(i)
// 	}
// 	return ret
// }
