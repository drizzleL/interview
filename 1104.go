package main

func pathInZigZagTree(label int) []int {
	if label == 0 {
		return nil
	}
	base := 1
	var flag bool
	for base*2 <= label {
		base *= 2
		flag = !flag
	}
	offset := label - base
	if flag {
		offset = base - offset - 1
	}
	var ret []int
	for base >= 1 {
		if !flag {
			ret = append(ret, base+offset)
		} else {
			ret = append(ret, base+(base-offset-1))
		}
		offset /= 2
		base /= 2
		flag = !flag
	}
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}
