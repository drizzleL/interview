package main

func findTheLongestBalancedSubstring(s string) int {
	var zeros, ones int
	var ret int
	for _, c := range s {
		switch c {
		case '0':
			if ones != 0 {
				ret = max(ret, min(zeros, ones)*2)
				ones = 0
				zeros = 0
			}
			zeros += 1
		case '1':
			ones += 1
		}
	}
	ret = max(ret, min(zeros, ones)*2)
	return ret
}
