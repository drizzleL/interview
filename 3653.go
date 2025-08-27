package main

func xorAfterQueries(nums []int, queries [][]int) int {
	for _, q := range queries {
		l, r := q[0], q[1]
		k, v := q[2], q[3]
		for i := l; i <= r; i += k {
			nums[i] *= v
			nums[i] %= 1e9 + 7
		}
	}
	var ret int
	for _, num := range nums {
		ret ^= num
	}
	return ret
}
