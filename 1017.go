package main

import "fmt"

func baseNeg2(n int) string {
	b1 := []byte(fmt.Sprintf("%b", n))
	rev := func(b []byte) {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}
	rev(b1)
	var b2 []byte
	var flag int
	for i := 0; i < len(b1); i++ {
		v := int(b1[i]-'0') + flag
		if v == 0 || v == 2 {
			b2 = append(b2, '0')
			continue
		}
		b2 = append(b2, '1')
		if i%2 == 1 {
			flag = 1
		}
	}
	if flag == 1 {
		b2 = append(b2, '1')
		if len(b2)%2 == 0 {
			b2 = append(b2, '1')
		}
	}
	rev(b2)
	return string(b2)
}
