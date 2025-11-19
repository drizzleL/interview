package main

func numberOfBeams(bank []string) int {
	var ret, last int
	for i := 0; i < len(bank); i++ {
		var cnt int
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == '1' {
				cnt += 1
			}
		}
		if cnt != 0 {
			ret += last * cnt
			last = cnt
		}
	}
	return ret
}
