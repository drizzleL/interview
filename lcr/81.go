package main

func combinationSum(candidates []int, target int) [][]int {
	var ret [][]int
	var helper func(i int, curr []int, tmp int)
	helper = func(i int, curr []int, tmp int) {
		if tmp == target {
			dst := make([]int, len(curr))
			copy(dst, curr)
			ret = append(ret, dst)
			return
		}
		if i == len(candidates) {
			return
		}
		helper(i+1, curr, tmp)
		for tmp < target {
			tmp += candidates[i]
			curr = append(curr, candidates[i])
			helper(i+1, curr, tmp)
		}
	}
	helper(0, nil, 0)
	return ret
}
