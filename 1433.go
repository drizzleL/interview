package main

func checkIfCanBreak(s1 string, s2 string) bool {
	var dict1, dict2 [26]int
	for i := range s1 {
		dict1[s1[i]-'a'] += 1
		dict2[s2[i]-'a'] += 1
	}
	check := func(a, b [26]int) bool {
		var sum int
		for i := 0; i < 26; i++ {
			sum += a[i]
			sum -= b[i]
			if sum < 0 {
				return false
			}
		}
		return true
	}
	return check(dict1, dict2) || check(dict2, dict1)
}
