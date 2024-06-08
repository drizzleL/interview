package main

// 0,0 = 0,0 + 0,1 + 0,2
// 0,1 = 0,0
// 0,2 = 0,1
// 1,0 = 1,0 + 1,1 + 1,2 + 0,1 + 0,2 + 0,0
// 1,1 = 1,0
// 1,2 = 1,1
func checkRecord(n int) int {
	if n == 0 {
		return 0
	}
	curr := [6]int{1, 1, 0, 1, 0, 0}
	mod := int(1e9 + 7)
	for i := 1; i < n; i++ {
		var next [6]int
		next[0] = (curr[0] + curr[1] + curr[2]) % mod
		next[1] = curr[0]
		next[2] = curr[1]
		next[3] = (curr[0] + curr[1] + curr[2] + curr[3] + curr[4] + curr[5]) % mod
		next[4] = curr[3]
		next[5] = curr[4]
		curr = next
	}
	var ret int
	for _, v := range curr {
		ret += v
		ret %= mod
	}
	return ret
}
