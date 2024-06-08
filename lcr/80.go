package main

func combine(n int, k int) [][]int {
	var ret [][]int
	var helper func(i int, tmp []int)
	helper = func(i int, tmp []int) {
		if len(tmp) == k {
			dst := make([]int, len(tmp))
			copy(dst, tmp)
			ret = append(ret, dst)
			return
		}
		if i > n {
			return
		}
		helper(i+1, append(tmp, i))
		helper(i+1, tmp)
	}
	helper(1, nil)
	return ret
}
