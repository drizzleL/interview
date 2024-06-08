package main

func pairSums(nums []int, target int) [][]int {
	dict := map[int]int{}
	var ret [][]int
	for _, num := range nums {
		another := target - num
		if dict[another] == 0 {
			dict[num]++
			continue
		}
		dict[another]--
		ret = append(ret, []int{num, another})
	}
	return ret
}
