package main

func shortestSuperstring(words []string) string {
	dict := map[[2]int]int{}
	calc := func(s1, s2 string) int {
		for size := len(s2); size > 0; size-- {
			if size > len(s1) {
				continue
			}
			if s1[len(s1)-size:] == s2[:size] {
				return size
			}
		}
		return 0
	}
	for i := range words {
		for j := range words {
			if i == j {
				continue
			}
			dict[[2]int{i, j}] = calc(words[i], words[j])
			dict[[2]int{j, i}] = calc(words[j], words[i])
		}
	}
	cache := map[[2]int]string{}
	var try func(from int, seen int) string
	try = func(from int, seen int) (ret string) {
		if seen+1 == 1<<len(words) {
			return words[from]
		}
		if c, ok := cache[[2]int{from, seen}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{from, seen}] = ret
		}()
		var w string
		if from != -1 {
			w = words[from]
		}
		for i := range words {
			if seen&(1<<i) != 0 {
				continue
			}
			tmp := try(i, seen|(1<<i))
			overlap := dict[[2]int{from, i}]
			tmp = w[:len(w)-overlap] + tmp
			if ret == "" || len(tmp) < len(ret) {
				ret = tmp
			}
		}
		return ret
	}
	return try(-1, 0)
}
