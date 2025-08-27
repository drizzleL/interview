package main

func possibleStringCount2(word string, k int) int {
	ret := 1
	var seg []int
	for i := 0; i < len(word); {
		c := word[i]
		var cnt int
		for i < len(word) && word[i] == c {
			i += 1
			cnt += 1
		}
		ret *= cnt
		ret %= 1e9 + 7
		k -= 1
		if cnt > 1 {
			seg = append(seg, cnt-1)
		}
	}
	if k <= 0 {
		return ret
	}
	dp := make([]int, k)
	dp[0] = 1
	for i := 0; i < len(seg); i++ {
		newDP := make([]int, k)
		var sum int = 0

		for j := 0; j < k; j++ {
			sum += dp[j]
			sum %= 1e9 + 7
			if j-seg[i]-1 >= 0 {
				sum -= dp[j-seg[i]-1]
				if sum < 0 {
					sum += 1e9 + 7
				}
			}
			newDP[j] = sum
		}
		dp = newDP
	}
	var ways int
	for i := range dp {
		ways += dp[i]
		ways %= 1e9 + 7
	}
	ret -= ways
	if ret < 0 {
		ret += 1e9 + 7
	}
	return ret
}
