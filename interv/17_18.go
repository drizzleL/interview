package main

func shortestSeq(big []int, small []int) []int {
	var numMatch int
	dict := map[int]int{}
	for _, val := range small {
		dict[val] = 0 // init candidate
	}
	m, n := 0, 0
	var ret []int
	for n < len(big) {
		for ; numMatch != len(small) && n < len(big); n++ {
			_, ok := dict[big[n]]
			if !ok {
				continue
			}
			dict[big[n]]++
			if dict[big[n]] == 1 { // first added
				numMatch++
			}
		}
		for ; numMatch == len(small); m++ {
			if len(ret) == 0 || n-1-m < ret[1]-ret[0] { // update
				ret = []int{m, n - 1}
			}
			_, ok := dict[big[m]]
			if !ok {
				continue
			}
			dict[big[m]]--
			if dict[big[m]] == 0 {
				numMatch--
			}
		}
	}
	return ret
}
