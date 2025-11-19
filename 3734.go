package main

func lexPalindromicPermutation(s string, target string) string {
	var dict [26]int
	for _, c := range s {
		dict[c-'a'] += 1
	}
	mid := -1
	for i, v := range dict {
		if v%2 == 1 {
			if mid != -1 {
				return ""
			}
			mid = i
		}
		dict[i] /= 2
	}
	maxB := make([]byte, len(s)/2)
	for i, j := 0, 25; i < len(maxB); j -= 1 {
		for k := 0; k < dict[j]; k++ {
			maxB[i] = 'a' + byte(j)
			i++
		}
	}
	helper := func(half []byte, mid int) string {
		halfSize := len(half)
		size := halfSize * 2
		if mid != -1 {
			size += 1
		}
		ret := make([]byte, size)
		copy(ret, half)
		if mid != -1 {
			ret[halfSize] = byte(mid + 'a')
		}
		for i, j := len(ret)-1, 0; j < halfSize; i, j = i-1, j+1 {
			ret[i] = half[j]
		}
		return string(ret)
	}
	if helper(maxB, mid) <= target {
		return ""
	}
	fill := func(b []byte, dict [26]int) []byte {
		for i := 0; i < 26; i++ {
			for dict[i] > 0 {
				b = append(b, byte('a'+i))
				dict[i] -= 1
			}
		}
		return b
	}
	tryFill := func(idx int, ret []byte) ([]byte, bool) {
		for j := idx + 1; j < 26; j++ {
			if dict[j] > 0 {
				dict[j] -= 1
				ret = append(ret, byte('a'+j))
				ret = fill(ret, dict)
				return ret, true
			}
		}
		return ret, false
	}
	getLarger := func(obj []byte) []byte {
		var ret []byte
		var ok bool
		for _, c := range obj {
			idx := int(c - 'a')
			if dict[idx] > 0 {
				dict[idx] -= 1
				ret = append(ret, c)
				continue
			}
			ret, ok = tryFill(idx, ret)
			if ok {
				return ret
			}
			for i := len(ret) - 1; i >= 0; i-- {
				c2 := int(ret[i] - 'a')
				dict[c2] += 1
				ret, ok = tryFill(c2, ret)
				if ok {
					return ret
				}
				ret = ret[:len(ret)-1]
			}
		}
		return ret
	}
	targetPre := []byte(target)[:len(target)/2]
	b := getLarger(targetPre)
	ret := helper(b, mid)
	if ret > target {
		return ret
	}
	reverse := func(b []byte) {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}
	nextPerm := func(b []byte) {
		for i := len(b) - 2; i >= 0; i-- {
			if b[i] >= b[i+1] {
				continue
			}
			reverse(b[i+1:])
			for m := i + 1; m < len(b); m++ {
				if b[m] > b[i] {
					b[m], b[i] = b[i], b[m]
					return
				}
			}
		}
	}
	nextPerm(b)
	return helper(b, mid)
}
