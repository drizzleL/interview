package main

func lenLongestFibSubseq(arr []int) int {
	var ret int
	dict := map[int]bool{}
	for _, num := range arr {
		dict[num] = true
	}
	vis := map[[2]int]bool{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			a, b := arr[i], arr[j]
			if vis[[2]int{a, b}] {
				continue
			}
			if !dict[a+b] {
				continue
			}
			curr := 2
			for dict[a+b] {
				vis[[2]int{a, b}] = true
				curr++
				a, b = b, a+b
			}
			if curr > ret {
				ret = curr
			}
		}
	}
	return ret
}

func lenLongestFibSubseq2(arr []int) int {
	var ret int
	dict := map[int]int{}
	for i, num := range arr {
		dict[num] = i
	}
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			dp[i][j] = 2
			lastVal := arr[j] - arr[i]
			kIdx, exists := dict[lastVal]
			if !exists {
				continue
			}
			if kIdx == i {
				continue
			}
			oldSize := dp[kIdx][i]
			if oldSize < 2 {
				oldSize = 2
			}
			if oldSize+1 > ret {
				ret = oldSize + 1
			}
			dp[i][j] = oldSize + 1
		}
	}
	return ret
}
