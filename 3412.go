package main

func calculateScore(s string) int64 {
	idxs := make([][]int, 26)
	var ret int
	for i, c := range s {
		v := int(c - 'a')
		mirror := (25 - v)
		if len(idxs[mirror]) == 0 {
			idxs[v] = append(idxs[v], i)
			continue
		}
		top := idxs[mirror][len(idxs[mirror])-1]
		ret += i - top
		idxs[mirror] = idxs[mirror][:len(idxs[mirror])-1]
	}
	return int64(ret)
}
