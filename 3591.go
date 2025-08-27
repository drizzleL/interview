package main

func checkPrimeFrequency(nums []int) bool {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num] += 1
	}
	isPrime := func(x int) bool {
		if x == 1 {
			return false
		}
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	for _, v := range dict {
		if isPrime(v) {
			return true
		}
	}
	return false
}
