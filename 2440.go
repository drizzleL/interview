package main

func componentValue(nums []int, edges [][]int) int {
	var sum int
	var maxVal int
	for _, num := range nums {
		sum += num
		maxVal = max(maxVal, num)
	}
	dict := make([][]int, len(nums))
	for _, ed := range edges {
		dict[ed[0]] = append(dict[ed[0]], ed[1])
		dict[ed[1]] = append(dict[ed[1]], ed[0])
	}
	var dfs func(x int, pre int, v int) int
	dfs = func(x int, pre int, v int) int {
		ret := nums[x]
		for _, next := range dict[x] {
			if next == pre {
				continue
			}
			tmp := dfs(next, x, v)
			if tmp == -1 {
				return -1
			}
			if tmp == v {
				continue
			}
			ret += tmp
		}
		if ret > v {
			return -1
		}
		return ret
	}
	for v := maxVal; v <= sum; v++ {
		if sum%v != 0 {
			continue
		}
		if dfs(0, -1, v) != v {
			continue
		}
		return sum/v - 1
	}
	return 0
}
