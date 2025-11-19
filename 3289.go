package main

func getSneakyNumbers(nums []int) []int {
	var ret []int
	seen := make([]bool, len(nums)-2)
	for _, num := range nums {
		if seen[num] {
			ret = append(ret, num)
			continue
		}
		seen[num] = true
	}
	return ret
}
