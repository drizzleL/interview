package main

func isTransformable(s string, t string) bool {
	var idxs [10][]int
	for i := range s {
		idxs[s[i]-'0'] = append(idxs[s[i]-'0'], i)
	}
	for i := 0; i < len(t); i++ {
		idx := int(t[i] - '0')
		if len(idxs[idx]) == 0 {
			return false
		}
		// pop idx
		pos := idxs[idx][0]
		idxs[idx] = idxs[idx][1:]
		for j := 0; j < idx; j++ {
			if len(idxs[j]) == 0 {
				continue
			}
			if idxs[j][0] < pos {
				return false
			}
		}
	}
	return true
}
