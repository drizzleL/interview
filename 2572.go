package main

func squareFreeSubsets(nums []int) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	dict := map[int]int{}
	for i, v := range primes {
		dict[v] = i
	}
	divide := func(x int) (int, bool) {
		var ret int
		for i := 0; x != 1 && i < len(primes); i++ {
			if x%primes[i] != 0 {
				continue
			}
			if x%(primes[i]*primes[i]) == 0 {
				return 0, false
			}
			ret |= 1 << i
		}
		return ret, true
	}
	valDict := map[int]int{
		0: 0,
	}
	vals := []int{0}
	cnts := []int{1}
	for _, num := range nums {
		arr, ok := divide(num)
		if !ok {
			continue
		}
		for _, val := range vals {
			if val&arr != 0 {
				continue
			}
			tmp := val | arr
			idx, ok := valDict[tmp]
			if !ok {
				idx = len(vals)
				valDict[tmp] = idx
				vals = append(vals, tmp)
				cnts = append(cnts, 0)
			}
			cnts[idx] += cnts[valDict[val]]
		}
	}
	var ret int
	for _, v := range cnts {
		ret += v
		ret %= 1e9 + 7
	}
	return ret - 1
}
