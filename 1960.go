package main

import "log"

func maxProduct(s string) int64 {
	paliLen := make([]int, len(s))
	var maxRight, mid int
	for i := 0; i < len(s); i++ {
		if i < maxRight {
			paliLen[i] = min(paliLen[mid*2-i], maxRight-i+1)
		} else {
			paliLen[i] = 1
		}
		for {
			l, r := i-paliLen[i], i+paliLen[i]
			if l < 0 || r >= len(s) {
				break
			}
			if s[l] != s[r] {
				break
			}
			paliLen[i] += 1
		}
		if i+paliLen[i]-1 > maxRight {
			maxRight = i + paliLen[i] - 1
			mid = i
		}
	}
	left, right := make([]int, len(s)), make([]int, len(s))
	maxRight = 0
	left[0] = 1
	for i := 1; i < len(s); i++ {
		if i > paliLen[maxRight]+maxRight-1 {
			maxRight = i
			left[i] = left[i-1]
		} else {
			left[i] = max(left[i-1], i-maxRight+1)
		}
	}
	maxLeft := len(s) - 1
	right[len(s)-1] = 1
	for i := len(s) - 2; i >= 0; i-- {
		if i < maxLeft+1-paliLen[maxLeft] {
			maxLeft = i
			right[i] = right[i+1]
		} else {
			right[i] = max(right[i+1], maxLeft-i+1)
		}
	}
	log.Println(left[0:100])
	var ret int
	for i := 0; i < len(s)-1; i++ {
		ret = max(ret, (left[i]*2-1)*(right[i+1]*2-1))
	}
	return int64(ret)
}
