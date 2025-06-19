package main

func smallestBeautifulString(s string, k int) string {
	b := []byte(s)
	i := len(s) - 1
	match := func(i int, c byte) bool { // can s[i] be c
		if i-1 >= 0 && b[i-1] == c {
			return false
		}
		if i-2 >= 0 && b[i-2] == c {
			return false
		}
		return true
	}
	var flag bool
	for i >= 0 && !flag {
		for j := s[i] + 1; j < 'a'+byte(k); j++ {
			if match(i, j) {
				b[i] = j
				flag = true
				break
			}
		}
		if flag {
			break
		}
		i--
	}
	if !flag {
		return ""
	}
	for i := i + 1; i < len(b); i++ {
		for j := 0; j < k; j++ {
			if match(i, 'a'+byte(j)) {
				b[i] = 'a' + byte(j)
				break
			}
		}
	}
	return string(b)
}
