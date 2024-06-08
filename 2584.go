package main

func findValidSplit(nums []int) int {
	dict := map[int]int{}
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	sieve := make([]bool, 1000)
	var primes []int
	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			continue
		}
		primes = append(primes, i)
		for j := 2; j*i < len(sieve); j++ {
			sieve[j*i] = true
		}
	}
	divide := func(x int) []int {
		var ret []int
		for _, p := range primes {
			if x%p != 0 {
				continue
			}
			ret = append(ret, p)
			for x != 1 && x%p == 0 {
				x /= p
			}
			if x == 1 {
				break
			}
		}
		if x != 1 {
			ret = append(ret, x)
		}
		return ret
	}
	for _, num := range nums {
		for _, v := range divide(num) {
			dict[v] += 1
		}
	}
	left := map[int]bool{}
	var used int
	for i := 0; i < len(nums)-1; i++ {
		num := nums[i]
		divs := divide(num)
		for _, div := range divs {
			left[div] = true
			dict[div] -= 1
			if dict[div] == 0 {
				used += 1
			}
		}
		if len(left) == used {
			return i
		}
	}
	return -1
}
