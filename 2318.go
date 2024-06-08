package main

func distinctSequences(n int) int {
	if n == 1 {
		return 6
	}
	conflicts := map[[2]int]bool{
		{1, 3}: true,
		{1, 5}: true,
		{2, 5}: true,
		{3, 1}: true,
		{3, 5}: true,
		{5, 3}: true,
		{5, 2}: true,
		{5, 1}: true,
	}
	dict := [36]int{}
	split := func(v int) (int, int) {
		return v / 6, v % 6
	}
	for k := range dict {
		a, b := split(k)
		if a == b {
			continue
		}
		if conflicts[[2]int{a, b}] {
			continue
		}
		dict[k] = 1
	}
	for i := 2; i < n; i++ {
		newDict := [36]int{}
		for j := 0; j < 6; j++ {
			for k, v := range dict {
				a, b := split(k)
				if a == j || b == j {
					continue
				}
				if conflicts[[2]int{j, a}] {
					continue
				}
				newDict[j*6+a] += v
				newDict[j*6+a] %= 1e9 + 7
			}
		}
		dict = newDict
	}
	var sum int
	for _, v := range dict {
		sum += v
		sum %= 1e9 + 7
	}
	return sum
}
