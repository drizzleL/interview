package main

func halvesAreAlike(s string) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	dict := map[byte]bool{}
	for _, v := range vowels {
		dict[v] = true
	}
	var cnt1, cnt2 int
	for i := 0; i < len(s)/2; i++ {
		if dict[s[i]] {
			cnt1 += 1
		}
	}
	for i := len(s) / 2; i < len(s); i++ {
		if dict[s[i]] {
			cnt2 += 1
		}
	}
	return cnt1 == cnt2
}
