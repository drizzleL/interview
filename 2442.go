package main

func countDistinctIntegers(nums []int) int {
	dict := map[int]bool{}
	reverse := func(x int) int {
		var ret int
		for x != 0 {
			ret = ret*10 + x%10
			x /= 10
		}
		return ret
	}
	for _, num := range nums {
		dict[num] = true
		dict[reverse(num)] = true
	}
	return len(dict)
}
