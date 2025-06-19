package main

func canMakeSubsequence(str1 string, str2 string) bool {
	for i, j := 0, 0; i < len(str1); i++ {
		a, b := str1[i]-'a', str2[j]-'a'
		if a == b || (a+1)%26 == b {
			j += 1
		}
		if j == len(str2) {
			return true
		}
	}
	return false
}
