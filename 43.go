package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var ret []int
	add := func(i int, v int) {
		if i == len(ret) {
			ret = append(ret, v)
		} else {
			ret[i] = v
		}
	}
	helper := func(idx int, a string, b int) {
		var carry int
		for i := len(a) - 1; i >= 0; i-- {
			v := int(a[i] - '0')
			v = v*b + carry
			if idx != len(ret) {
				v += ret[idx]
			}

			carry = v / 10
			v %= 10
			add(idx, v)
			idx += 1
		}
		if carry != 0 {
			add(idx, carry)
		}
	}
	for i, start := len(num2)-1, 0; i >= 0; i, start = i-1, start+1 {
		v2 := int(num2[i] - '0')
		helper(start, num1, v2)
	}
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	b := make([]byte, 0, len(ret))
	for _, c := range ret {
		b = append(b, byte('0'+c))
	}
	return string(b)
}
