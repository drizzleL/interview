package main

func maxCoins6(nums []int) int {
	newNums := make([]int, 0, len(nums)+2)
	newNums = append(newNums, 1)
	for _, num := range nums {
		if num == 0 {
			continue
		}
		newNums = append(newNums, num)
	}
	newNums = append(newNums, 1)
	cache := make([][]int, len(newNums))
	for i := range cache {
		cache[i] = make([]int, len(newNums))
	}
	var helper func(l, r int) int
	helper = func(l, r int) (ret int) {
		if l+1 == r {
			return 0
		}
		if cache[l][r] != 0 {
			return cache[l][r]
		}
		defer func() {
			cache[l][r] = ret
		}()
		for i := l + 1; i < r; i++ {
			ret = max(ret, helper(l, i)+helper(i, r)+newNums[l]*newNums[i]*newNums[r])
		}
		return
	}
	return helper(0, len(newNums)-1)
}
