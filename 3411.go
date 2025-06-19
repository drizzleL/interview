package main

func maxLength(nums []int) int {
	primes := []int{2, 3, 5, 7}
	split := func(x int) []int {
		var ret []int
		for _, v := range primes {
			if x%v == 0 {
				ret = append(ret, v)
			}
		}
		return ret
	}
	dict := map[int]int{}
	check := func() bool {
		for _, v := range dict {
			if v > 1 {
				return false
			}
		}
		return true
	}
	ret := 2
	for l, r := 0, 0; r < len(nums); r++ {
		for _, v := range split(nums[r]) {
			dict[v] += 1
		}
		for !check() {
			for _, v := range split(nums[l]) {
				dict[v] -= 1
			}
			l++
		}
		ret = max(ret, r-l+1)
	}
	return ret
}
