package main

func breakPalindrome(palindrome string) string {
	if len(palindrome) <= 1 {
		return ""
	}
	ret := []byte(palindrome)
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		if ret[i] != 'a' {
			ret[i] = 'a'
			return string(ret)
		}
	}
	ret[len(ret)-1] = 'b'
	return string(ret)
}
