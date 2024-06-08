package main

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	dict := make([][]int, 26)
	for i := range dict {
		dict[i] = make([]int, 26)
		for j := range dict[i] {
			dict[i][j] = -1
		}
	}
	for i := 0; i < 26; i++ {
		dict[i][i] = 0
	}
	for idx := range original {
		i, j := int(original[idx]-'a'), int(changed[idx]-'a')
		c := cost[idx]
		for a := 0; a < 26; a++ {
			for b := 0; b < 26; b++ {
				if dict[a][i] == -1 || dict[j][b] == -1 {
					continue
				}
				if dict[a][b] == -1 || dict[a][b] > dict[a][i]+c+dict[j][b] {
					dict[a][b] = dict[a][i] + c + dict[j][b]
				}
			}
		}
	}

	var ret int
	for idx := range source {
		i, j := int(source[idx]-'a'), int(target[idx]-'a')
		if dict[i][j] == -1 {
			return -1
		}
		ret += dict[i][j]
	}
	return int64(ret)
}
