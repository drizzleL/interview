package main

func minimumPartition(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	var tmp int
	ret := 1
	for _, c := range s {
		d := int(c - '0')
		if d > k { // early return
			return -1
		}
		if tmp+d <= k {
			tmp += d
		} else {
			ret += 1
			tmp = d
		}
		tmp *= 10
	}
	return ret
}
