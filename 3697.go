package main

func decimalRepresentation(n int) []int {
	var ret []int
	for base := 1; n > 0; n, base = n/10, base*10 {
		if n%10 == 0 {
			continue
		}
		ret = append(ret, n%10*base)
	}
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}
