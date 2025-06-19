package main

func isAdditiveNumber(num string) bool {
	reverse := func(b []byte) {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}
	add := func(b1, b2 []byte) []byte {
		reverse(b1)
		reverse(b2)
		var b []byte
		var flag int
		for i := 0; i < max(len(b1), len(b2)); i++ {
			if i < len(b1) {
				flag += int(b1[i] - '0')
			}
			if i < len(b2) {
				flag += int(b2[i] - '0')
			}
			b = append(b, byte(flag%10)+'0')
			flag /= 10
		}
		if flag != 0 {
			b = append(b, 1+'0')
		}
		reverse(b)
		return b
	}
	equal := func(i, j, k int) int {
		b1, b2 := []byte(num[i:j+1]), []byte(num[j+1:k+1])
		sum := add(b1, b2)
		if len(num)-k-1 < len(sum) {
			return -1
		}
		for m := 0; m < len(sum); m++ {
			if sum[m] != num[k+1+m] {
				return -1
			}
		}
		return k + len(sum)
	}
	check := func(i, j, k int) bool {
		if num[i] == '0' && i != j {
			return false
		}
		if num[j+1] == '0' && j+1 != k {
			return false
		}
		for {
			end := equal(i, j, k)
			if end == -1 {
				return false
			}
			if end == len(num)-1 {
				return true
			}
			i, j, k = j+1, k, end
		}
	}
	for j := 0; j < len(num); j++ {
		for k := j + 1; k < len(num); k++ {
			if check(0, j, k) {
				return true
			}
		}
	}
	return false
}
