package main

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	dict := make([][]int, len(edges)+1)
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	seen := make([]bool, len(edges)+1)
	seen[0] = true
	var helper func(x int) (sum, val int)
	helper = func(x int) (sum, val int) {
		for _, next := range dict[x] {
			if seen[next] {
				continue
			}
			seen[next] = true
			childSum, val2 := helper(next)
			sum += childSum
			val += val2
		}
		if sum == 0 {
			return values[x], 0
		}
		return sum + values[x], max(sum, val+values[x])
	}
	sum, val := helper(0)
	return int64(max(sum-values[0], val))
}
