package main

func subsets(nums []int) [][]int {
	var ret [][]int
	ret = append(ret, []int{})
	for _, num := range nums {
		size := len(ret)
		for i := 0; i < size; i++ {
			tmp := make([]int, len(ret[i]), len(ret[i])+1)
			copy(tmp, ret[i])
			tmp = append(tmp, num)
			ret = append(ret, tmp)
		}
	}
	return ret
}
