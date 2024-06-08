package main

func longestIdealString(s string, k int) int {
	var dict [26]int
	var ret int
	for _, c := range s {
		d := int(c - 'a')
		var tmp int
		for i := range dict {
			if abs(i-d) > k {
				continue
			}
			tmp = max(tmp, dict[i]+1)
		}
		dict[d] = tmp
		ret = max(ret, tmp)
	}
	return ret
}

func secondsToRemoveOccurrences(s string) int {
	var zeros int
	var ret int
	for _, c := range s {
		if c == '0' {
			zeros += 1
		} else if zeros > 0 {
			ret = max(ret+1, zeros)
		}
	}
	return ret
}
