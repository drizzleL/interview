package main

func scoreOfStudents(s string, answers []int) int {
	nums := []int{int(s[0] - '0')}
	for i := 1; i < len(s); i += 2 {
		val := int(s[i+1] - '0')
		switch s[i] {
		case '+':
			nums = append(nums, val)
		case '*':
			nums[len(nums)-1] *= val
		}
	}
	var correct int
	for _, num := range nums {
		correct += num
	}
	cache := map[[2]int]map[int]bool{}
	var helper func(l, r int) (ret map[int]bool)
	helper = func(l, r int) (ret map[int]bool) {
		if c, ok := cache[[2]int{l, r}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{l, r}] = ret
		}()
		dict := map[int]bool{}
		if l == r {
			dict[int(s[l]-'0')] = true
			return dict
		}
		for i := l + 1; i < r; i += 2 {
			ldict, rdict := helper(l, i-1), helper(i+1, r)
			switch s[i] {
			case '*':
				for v := range ldict {
					for v2 := range rdict {
						tmp := v * v2
						if tmp > 1000 {
							continue
						}
						dict[tmp] = true
					}
				}
			case '+':
				for v := range ldict {
					for v2 := range rdict {
						tmp := v + v2
						if tmp > 1000 {
							continue
						}
						dict[tmp] = true
					}
				}
			}
		}
		return dict
	}
	dict := helper(0, len(s)-1)
	var ret int
	for _, ans := range answers {
		switch {
		case ans == correct:
			ret += 5
		case dict[ans]:
			ret += 2
		}
	}
	return ret
}
