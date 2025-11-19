package main

func hasSameDigits(s string) bool {
	var vals []int
	for _, c := range s {
		vals = append(vals, int(c-'0'))
	}
	for len(vals) != 2 {
		for i := 0; i < len(vals)-1; i++ {
			vals[i] = (vals[i] + vals[i+1]) % 10
		}
		vals = vals[:len(vals)-1]
	}
	return vals[0] == vals[1]
}
