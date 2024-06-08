package main

func equalSubstring(s string, t string, maxCost int) int {
	leftCost := maxCost
	var ret int
	var l, r int
	for r < len(t) {
		for r < len(t) && leftCost >= 0 {
			leftCost -= abs(int(s[r]) - int(t[r]))
			ret = max(ret, r-l)
			r += 1
		}
		for l < r && leftCost < 0 {
			leftCost += abs(int(s[l]) - int(t[l]))
			l += 1
		}
		ret = max(ret, r-l)
	}
	return ret
}
