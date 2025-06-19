package main

func combinationSum3(k int, n int) [][]int {
	var ret [][]int
	var helper func(i int, tmp []int, sum int)
	helper = func(i int, tmp []int, sum int) {
		if len(tmp) == k {
			if sum == n {
				cpy := make([]int, len(tmp))
				copy(cpy, tmp)
				ret = append(ret, cpy)
			}
			return
		}
		if sum >= n {
			return
		}
		if i == 10 {
			return
		}
		helper(i+1, append(tmp, i), sum+i)
		helper(i+1, tmp, sum)
	}
	helper(1, nil, 0)
	return ret
}
