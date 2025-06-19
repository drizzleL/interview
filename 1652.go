package main

func decrypt(code []int, k int) []int {
	if k == 0 {
		return make([]int, len(code))
	}
	sums := make([]int, 1, len(code))
	for i := range code {
		sums = append(sums, sums[len(sums)-1]+code[i])
	}
	getCode := func(l, r int) int {
		var ret int
		if l >= len(code) {
			ret = sums[r%len(code)+1] - sums[l%len(code)]
		} else if r < 0 {
			ret = sums[r+len(code)+1] - sums[l+len(code)]
		} else if l < 0 {
			ret = sums[r+1] + sums[len(code)+1] - sums[l+len(code)]
		} else if r >= len(code) {
			ret = sums[len(code)+1] - sums[l] + sums[r-len(code)+1]
		} else {
			ret = sums[r+1] - sums[l]
		}
		return ret
	}
	ret := make([]int, len(code))
	for i := range code {
		var l, r int
		if k > 0 {
			l, r = i+1, i+k
		} else {
			l, r = i+k, i-1
		}
		ret[i] = getCode(l, r)
	}
	return ret
}
