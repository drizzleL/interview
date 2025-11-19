package main

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	known := make([][]bool, len(languages))
	for i := range known {
		known[i] = make([]bool, n)
	}
	for user, lang := range languages {
		for _, l := range lang {
			known[user][l-1] = true
		}
	}
	var fr2 [][]int
	for _, f := range friendships {
		a, b := f[0]-1, f[1]-1
		var flag bool
		for _, l := range languages[a] {
			if known[b][l-1] {
				flag = true
				break
			}
		}
		if !flag {
			fr2 = append(fr2, []int{a, b})
		}
	}
	if len(fr2) == 0 {
		return 0
	}
	ret := len(languages)
	for i := 0; i < n; i++ {
		dict := map[int]bool{}
		for _, f := range fr2 {
			a, b := f[0], f[1]
			k1, k2 := known[a][i], known[b][i]
			if !k1 {
				dict[a] = true
			}
			if !k2 {
				dict[b] = true
			}
		}
		ret = min(ret, len(dict))
	}
	return ret
}
