package main

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	size := len(students)
	comp := func(a, b []int) int {
		var ret int
		for i := 0; i < len(a); i++ {
			if a[i] == b[i] {
				ret += 1
			}
		}
		return ret
	}
	cache := map[[2]int]int{}
	var helper func(i int, mask int) int
	helper = func(i int, mask int) (ret int) {
		if i == size {
			return 0
		}
		if c, ok := cache[[2]int{i, mask}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{i, mask}] = ret
		}()
		for j := 0; j < size; j++ {
			if (1<<j)&mask != 0 {
				continue
			}
			ret = max(ret, comp(students[i], mentors[j])+helper(i+1, mask|(1<<j)))
		}
		return ret
	}
	return helper(0, 0)
}
