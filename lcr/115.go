package main

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	in := make([]int, len(nums)+1)
	dict := make([][]int, len(nums)+1)
	for _, seq := range sequences {
		for i := 1; i < len(seq); i++ {
			in[seq[i]]++
			dict[seq[i-1]] = append(dict[seq[i-1]], seq[i])
		}
	}
	var v int
	for i := 1; i < len(in); i++ {
		if in[i] == 0 {
			if v != 0 { // multiple zero in element
				return false
			}
			v = i
		}
	}
	for v != 0 {
		var nextv int
		for _, next := range dict[v] {
			in[next]--
			if in[next] == 0 {
				if nextv != 0 {
					return false
				}
				nextv = next
			}
		}
		v = nextv
	}
	return true
}
