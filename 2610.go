package main

func findMatrix(nums []int) [][]int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num] += 1
	}
	var ret [][]int
	for len(dict) != 0 {
		var tmp []int
		for k, v := range dict {
			tmp = append(tmp, k)
			if v == 1 {
				delete(dict, k)
				continue
			}
			dict[k] -= 1
		}
		ret = append(ret, tmp)
	}
	return ret
}
