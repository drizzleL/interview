package main

func minSplitMerge(nums1 []int, nums2 []int) int {
	size := len(nums1)
	var v1, v2 [6]int
	for i, num := range nums1 {
		v1[i] = num
		v2[i] = nums2[i]
	}
	cache := map[[6]int]bool{}
	cache[v1] = true
	nodes := [][6]int{v1}
	insert := func(v1 [6]int, l, r int, start, end int) [][6]int {
		a := v1[start : end+1]
		b := append([]int{}, v1[l:start]...)
		b = append(b, v1[end+1:r+1]...)
		var ret [][6]int
		for insert := 0; insert <= len(b); insert++ {
			c := append([]int{}, b[:insert]...)
			c = append(c, a...)
			c = append(c, b[insert:]...)
			tmp := v1
			for i := 0; i < len(c); i++ {
				tmp[l+i] = c[i]
			}
			ret = append(ret, tmp)
		}
		return ret
	}
	getNext := func(v1 [6]int) [][6]int {
		var ret [][6]int
		l, r := 0, size-1
		for v1[l] == v2[l] {
			l++
		}
		for v1[r] == v2[r] {
			r--
		}
		for splitSize := 1; splitSize <= r-l; splitSize++ {
			for start := l; start <= r-splitSize+1; start++ {
				end := start + splitSize - 1
				ret = append(ret, insert(v1, l, r, start, end)...)
			}
		}
		return ret
	}
	for step := 0; len(nodes) != 0; step++ {
		var nextNodes [][6]int
		for _, node := range nodes {
			if node == v2 {
				return step
			}
			for _, next := range getNext(node) {
				if _, ok := cache[next]; ok {
					continue
				}
				cache[next] = true
				nextNodes = append(nextNodes, next)
			}
		}
		nodes = nextNodes
	}
	return -1
}
