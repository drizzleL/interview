package main

func rearrangeArray(nums []int) []int {
	var ret []int
	var pos, neg []int
	for _, num := range nums {
		if num > 0 {
			pos = append(pos, num)
		} else {
			neg = append(neg, num)
		}
	}
	for i := range pos {
		ret = append(ret, pos[i], neg[i])
	}
	return ret
}
