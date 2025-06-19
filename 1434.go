package main

func numberWays(hats [][]int) int {
	var hatToPeople [40][]int
	for i, h := range hats {
		for _, p := range h {
			hatToPeople[p] = append(hatToPeople[p], i)
		}
	}
	cache := map[[2]int]int{}
	var wear func(dict int, i int) int
	wear = func(dict int, i int) (ret int) {
		if dict == (1<<len(hats))-1 {
			return 1
		}
		if i == 40 {
			return 0
		}
		if c, ok := cache[[2]int{dict, i}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{dict, i}] = ret
		}()
		ret += wear(dict, i+1)
		for _, p := range hatToPeople[i] {
			if dict&(1<<p) != 0 {
				continue
			}
			ret += wear(dict|(1<<p), i+1)
		}
		return
	}
	return wear(0, 0)
}
