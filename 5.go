package main

func longestPalindrome2(s string) string {
	b := make([]byte, len(s)*2+1)
	b[0] = '#'
	for i := range s {
		b[2*i+1] = s[i]
		b[2*i+2] = '#'
	}
	paliLen := make([]int, len(b))
	var maxRight, mid int
	paliLen[0] = 1
	for i := 1; i < len(b); i++ {
		if i <= maxRight {
			paliLen[i] = min(paliLen[mid*2-i], maxRight-i+1)
		} else {
			paliLen[i] = 1
		}
		l, r := i-paliLen[i], i+paliLen[i]
		for l >= 0 && r < len(b) {
			if b[l] != b[r] {
				break
			}
			paliLen[i] += 1
			l -= 1
			r += 1
		}
		if i+paliLen[i]-1 > maxRight {
			maxRight = i + paliLen[i] - 1
			mid = i
		}
	}
	var maxSize, midIdx int
	getSize := func(i int) int {
		switch b[i] {
		case '#':
			return paliLen[i] / 2 * 2
		default:
			return (paliLen[i]/2-1)*2 + 1
		}
	}
	for i := range paliLen {
		size := getSize(i)
		if size > maxSize {
			maxSize = size
			midIdx = i
		}
	}
	ret := make([]byte, 0, maxSize)
	for i := midIdx - paliLen[midIdx] + 1; i < midIdx+paliLen[midIdx]; i++ {
		if b[i] == '#' {
			continue
		}
		ret = append(ret, b[i])
	}
	return string(ret)
}
