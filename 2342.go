package main

func maximumSum(nums []int) int {
	dict := map[int]int{}
	ret := -1
	for _, num := range nums {
		var tmp int
		for num2 := num; num2 != 0; num2 /= 10 {
			tmp += num % 10
		}
		if last, ok := dict[tmp]; ok {
			ret = max(ret, last+num)
		}
		dict[tmp] = max(dict[tmp], num)
	}
	return ret
}
