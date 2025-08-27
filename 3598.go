package main

func longestCommonPrefix(words []string) []int {
	helper := func(a, b string) int {
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] != b[i] {
				return i
			}
		}
		return min(len(a), len(b))
	}
	left := make([]int, len(words))
	right := make([]int, len(words))
	for i := 1; i < len(words); i++ {
		left[i] = max(left[i-1], helper(words[i-1], words[i]))
	}
	for i := len(words) - 1; i >= 0; i-- {
		right[i] = max(right[i+1], helper(words[i], words[i+1]))
	}
	ret := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		if i != 0 && i != len(words)-1 {
			ret[i] = helper(words[i-1], words[i+1])
		}
		if i != 0 {
			ret[i] = max(ret[i], left[i-1])
		}
		if i != len(words)-1 {
			ret[i] = max(ret[i], right[i+1])
		}
	}
	return ret
}
