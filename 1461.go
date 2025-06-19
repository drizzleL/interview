package main

func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}
	mask := 1<<k - 1
	var num int
	for i := 0; i < k-1; i++ {
		num <<= 1
		num |= int(s[i] - '0')
	}
	seen := make([]bool, 1<<k)
	for i := k - 1; i < len(s); i++ {
		num <<= 1
		num |= int(s[i] - '0')
		num &= mask
		seen[num] = true
	}
	for _, v := range seen {
		if !v {
			return false
		}
	}
	return true
}
