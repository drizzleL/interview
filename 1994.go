package main

func numberOfGoodSubsets(nums []int) int {
	primes := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
	}
	mask := func(x int) int {
		var ret int
		for i := 0; i < len(primes); i++ {
			if x%primes[i] != 0 {
				continue
			}
			x /= primes[i]
			if x%primes[i] == 0 {
				return -1
			}
			ret |= 1 << i
		}
		return ret
	}
	maskDict := make([]int, 31)
	for i := 1; i <= 30; i++ {
		maskDict[i] = mask(i)
	}
	dict := make([]int, 1024)
	dict[0] = 1
	var ret int
	cntDict := map[int]int{}
	for _, num := range nums {
		cntDict[num] += 1
	}
	for num, cnt := range cntDict {
		if num == 1 {
			continue
		}
		m := maskDict[num]
		if m == -1 {
			continue
		}
		for k, v := range dict {
			if m&k != 0 {
				continue
			}
			if v == 0 {
				continue
			}
			dict[m|k] += v * cnt
			dict[m|k] %= 1e9 + 7
		}
	}
	for num, v := range dict {
		if num == 0 {
			continue
		}
		ret += v
		ret %= 1e9 + 7
	}
	for i := 0; i < cntDict[1]; i++ {
		ret *= 2
		ret %= 1e9 + 7
	}
	return ret
}
