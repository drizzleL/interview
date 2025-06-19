package main

func getHappyString(n int, k int) string {
	b := make([]byte, 0, n)
	cnt := 1
	for i := 1; i < n; i++ {
		cnt *= 2
	}
	if k > cnt*3 {
		return ""
	}
	k -= 1
	var pre int
	if k >= cnt*2 {
		pre = 2
	} else if k >= cnt {
		pre = 1
	} else {
		pre = 0
	}
	cc := []byte{'a', 'b', 'c'}
	b = append(b, cc[pre])
	getNext := func(pre int) int {
		var pres []int
		switch pre {
		case 0:
			pres = []int{1, 2}
		case 1:
			pres = []int{0, 2}
		case 2:
			pres = []int{0, 1}
		}
		if k < cnt {
			pre = pres[0]
		} else {
			pre = pres[1]
		}
		pre = (pre + 3) % 3
		return pre
	}
	for len(b) < n {
		k %= cnt
		cnt /= 2
		pre = getNext(pre)
		b = append(b, cc[pre])
	}
	return string(b)
}
