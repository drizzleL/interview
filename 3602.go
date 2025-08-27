package main

func concatHex36(n int) string {
	v1, v2 := n*n, n*n*n
	var b []byte
	to := func(v int) byte {
		if v < 10 {
			return byte(v + '0')
		}
		return byte(v - 10 + 'A')
	}
	trans := func(v int, unit int) []byte {
		var ret []byte
		for base := 1; v != 0; v, base = v/unit, base*unit {
			ret = append(ret, to(v%unit))
		}
		for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
			ret[i], ret[j] = ret[j], ret[i]
		}
		return ret
	}
	b = append(b, trans(v1, 16)...)
	b = append(b, trans(v2, 36)...)
	return string(b)
}
