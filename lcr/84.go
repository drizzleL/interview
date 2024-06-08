package main

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	dict := map[int]int{}
	for _, num := range nums {
		dict[num]++
	}
	var makeArrs func(num []int, curr []int, i int, v int, c int, arr *[][]int)
	makeArrs = func(num []int, curr []int, i int, v int, c int, arr *[][]int) {
		if i == len(num) {
			for m := 0; m < c; m++ {
				curr = append(curr, v)
			}
			*arr = append(*arr, append([]int{}, curr...))
			return
		}
		if c == 0 {
			curr = append(curr, num[i:]...)
			*arr = append(*arr, append([]int{}, curr...))
			return
		}
		makeArrs(num, append(curr, num[i]), i+1, v, c, arr)
		for m := 0; m < c; m++ {
			curr = append(curr, v)
			makeArrs(num, append(curr, num[i]), i+1, v, c-m-1, arr)
		}
	}
	var tmp []int
	for m := 0; m < dict[nums[0]]; m++ {
		tmp = append(tmp, nums[0])
	}
	ret := [][]int{tmp}
	insert := func(nums []int, v int) [][]int {
		c := dict[v]
		var arr [][]int
		makeArrs(nums, nil, 0, v, c, &arr)
		return arr
	}
	helper := func(v int) {
		size := len(ret)
		var tmp [][]int
		for i := 0; i < size; i++ {
			tmp = append(tmp, insert(ret[i], v)...)
		}
		ret = tmp
	}
	for k := range dict {
		if k == nums[0] {
			continue
		}
		helper(k)
	}
	return ret
}
