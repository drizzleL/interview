package main

func minJumps(arr []int) int {
	arr = func() []int {
		var after []int
		for i, v := range arr {
			if i >= 2 && arr[i-1] == arr[i] && arr[i-2] == arr[i] {
				continue
			}
			after = append(after, v)
		}
		return after
	}()
	dict := map[int][]int{}
	for i, v := range arr {
		dict[v] = append(dict[v], i)
	}
	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = -1
	}
	nodes := []int{0}
	dp[0] = 0
	for i := 0; dp[len(arr)-1] == -1; i++ {
		var next []int
		for _, n := range nodes {
			if n > 0 && dp[n-1] == -1 {
				dp[n-1] = i + 1
				next = append(next, n-1)
			}
			if n < len(arr)-1 && dp[n+1] == -1 {
				dp[n+1] = i + 1
				next = append(next, n+1)
			}
			v := arr[n]
			for _, j := range dict[v] {
				if dp[j] == -1 {
					dp[j] = i + 1
					next = append(next, j)
				}
			}
		}
		nodes = next
	}
	return dp[len(arr)-1]
}
