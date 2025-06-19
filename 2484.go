package main

func countPalindromes(s string) int {
	var left, right [100]int
	var ret int
	mirror := func(i int) int {
		return i%10*10 + i/10
	}
	var leftCnt, rightCnt [10]int
	for i := 0; i < 2; i++ {
		d := int(s[i] - '0')
		if i != 0 {
			for j := 0; j <= 9; j++ {
				left[j*10+d] += leftCnt[j]
				left[j*10+d] %= 1e9 + 7
			}
		}
		leftCnt[d] += 1
		leftCnt[d] %= 1e9 + 7
	}
	for i := len(s) - 1; i > 2; i-- {
		d := int(s[i] - '0')
		if i != len(s)-1 {
			for j := 0; j <= 9; j++ {
				right[d*10+j] += rightCnt[j]
				right[d*10+j] %= 1e9 + 7
			}
		}
		rightCnt[d] += 1
		rightCnt[d] %= 1e9 + 7
	}
	for mid := 2; mid < len(s)-2; mid++ {
		for i := 0; i <= 99; i++ {
			j := mirror(i)
			ret += left[i] * right[j]
			ret %= 1e9 + 7
		}
		d := int(s[mid] - '0')
		for i := 0; i <= 9; i++ {
			left[i*10+d] += leftCnt[i]
			left[i*10+d] %= 1e9 + 7
		}
		leftCnt[d]++
		d2 := int(s[mid+1] - '0')
		rightCnt[d2]--
		for i := 0; i <= 9; i++ {
			right[d2*10+i] -= rightCnt[i]
			if right[d2*10+i] < 0 {
				right[d2*10+i] += 1e9 + 7
			}
		}
	}
	return ret
}
