package main

func subarraysDivByK(nums []int, k int) int {
	var sum int
	dict := map[int]int{
		0: 1,
	}
	var ret int
	for _, num := range nums {
		sum += num
		sum %= k
		if sum < 0 {
			sum += k
		}
		ret += dict[sum]
		dict[sum] += 1
	}
	return ret
}
