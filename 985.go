package main

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	var sum int
	for _, num := range nums {
		if num%2 != 0 {
			continue
		}
		sum += num
	}
	ret := make([]int, 0, len(queries))
	for _, q := range queries {
		v, i := q[0], q[1]
		if nums[i]%2 == 0 { // even before
			sum -= nums[i]
		}
		nums[i] += v
		if nums[i]%2 == 0 {
			sum += nums[i]
		}
		ret = append(ret, sum)
	}
	return ret
}
