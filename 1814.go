package main

func countNicePairs(nums []int) int {
	dict := map[int]int{}
	var ret int
	rev := func(x int) int {
		var ret int
		for x != 0 {
			ret = ret*10 + x%10
			x /= 10
		}
		return ret
	}
	for _, num := range nums {
		tmp := num - rev(num)
		ret += dict[tmp]
		dict[tmp] += 1
		ret %= 1e9 + 7
	}
	return ret
}
