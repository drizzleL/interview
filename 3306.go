package main

func countOfSubstrings(word string, k int) int64 {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	vowelBit := func(b byte) int {
		for i := range vowels {
			if vowels[i] == b {
				return i
			}
		}
		return -1
	}
	isCons := func(r byte) bool {
		return vowelBit(r) == -1
	}
	nextCons := make([]int, len(word))
	lastCons := len(word)
	for i := len(word) - 1; i >= 0; i-- {
		nextCons[i] = lastCons
		if isCons(word[i]) {
			lastCons = i
		}
	}
	var ret int
	var consCnt int
	var dict [5]int
	vowelMatch := func() bool {
		for _, v := range dict {
			if v == 0 {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < len(word); r++ {
		bit := vowelBit(word[r])
		if bit == -1 {
			consCnt += 1
		} else {
			dict[bit] += 1
		}
		for consCnt > k {
			lBit := vowelBit(word[l])
			if lBit == -1 {
				consCnt -= 1
			} else {
				dict[lBit] -= 1
			}
			l++
		}
		for l < r && vowelMatch() && consCnt == k {
			ret += nextCons[r] - r
			lBit := vowelBit(word[l])
			if lBit == -1 {
				consCnt -= 1
			} else {
				dict[lBit] -= 1
			}
			l++
		}
	}
	return int64(ret)
}
