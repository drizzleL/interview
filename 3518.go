package main

func smallestPalindrome2(s string, k int) string {
	var dict [26]int
	for _, c := range s {
		dict[c-'a'] += 1
	}
	b := make([]byte, len(s))
	for k, v := range dict {
		if v%2 == 1 {
			b[len(s)/2] = byte('a' + k)
			dict[k] -= 1
		}
		dict[k] /= 2
	}
	var combination func(a int, b int, k int) int
	combination = func(a int, b int, k int) int {
		if a < b {
			return 0
		}
		if b*2 > a {
			return combination(a, a-b, k)
		}
		ret := 1
		for i := 0; i < b; i++ {
			ret *= (a - i)
			ret /= i + 1
			if ret > k {
				return k + 1
			}
		}
		return ret
	}
	comb := func(size int, k int) int {
		ret := 1
		for _, v := range dict {
			ret *= combination(size, v, k)
			if ret > k {
				return ret
			}
			size -= v
		}
		return ret
	}
	if comb(len(s)/2, k) < k {
		return ""
	}
	var helper func(i int, k int)
	helper = func(i int, k int) {
		if i == len(s)/2 {
			return
		}
		for j, v := range dict {
			if v == 0 {
				continue
			}
			dict[j] -= 1
			c := comb(len(s)/2-i-1, k)
			if c >= k {
				b[i] = byte('a' + j)
				helper(i+1, k)
				break
			}
			dict[j] += 1
			k -= c
		}
	}
	helper(0, k)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b[j] = b[i]
	}
	return string(b)
}
