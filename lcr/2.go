package main

func addBinary(a string, b string) string {
	size := len(a)
	if len(b) > size {
		size = len(b)
	}
	bb := make([]byte, size+1)
	var idx int
	var carry byte
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		for i >= 0 && j >= 0 {
			tmp := carry + a[i] - '0' + b[j] - '0'
			carry, bb[idx] = tmp/2, tmp%2+'0'
			idx += 1
			i -= 1
			j -= 1
		}
		for i >= 0 {
			tmp := carry + a[i] - '0'
			carry, bb[idx] = tmp/2, tmp%2+'0'
			idx += 1
			i -= 1
		}
		for j >= 0 {
			tmp := carry + b[j] - '0'
			carry, bb[idx] = tmp/2, tmp%2+'0'
			idx += 1
			j -= 1
		}
	}
	if carry == 1 {
		bb[len(bb)-1] = '1'
	} else {
		bb = bb[:len(bb)-1]
	}
	for i, j := 0, len(bb)-1; i < j; i, j = i+1, j-1 {
		bb[i], bb[j] = bb[j], bb[i]
	}
	return string(bb)
}
