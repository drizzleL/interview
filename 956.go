package main

func tallestBillboard(rods []int) int {
	dict := map[int]int{
		0: 0,
	}
	for i := 0; i < len(rods); i++ {
		dictClone := map[int]int{}
		for k, v := range dict {
			dictClone[k] = v
		}
		for k, v := range dictClone {
			dict[k+rods[i]] = max(dict[k+rods[i]], v)
			dict[abs(k-rods[i])] = max(dict[abs(k-rods[i])], v+min(v, rods[i]))
		}
	}
	return dict[0]
}
