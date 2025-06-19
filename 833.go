package main

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	replaced := make([]int, len(s))
	for i := range replaced {
		replaced[i] = -1
	}
	check := func(i, idx int) bool {
		if i+len(sources[idx]) >= len(s)+1 {
			return false
		}
		return sources[idx] == s[i:i+len(sources[idx])]
	}
	for idx, i := range indices {
		if check(i, idx) {
			replaced[i] = idx
		}
	}
	var ret []byte
	for i := 0; i < len(s); {
		if replaced[i] != -1 {
			idx := replaced[i]
			i += len(sources[idx])
			ret = append(ret, []byte(targets[idx])...)
			continue
		}
		ret = append(ret, s[i])
		i++
	}
	return string(ret)
}
