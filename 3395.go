package main

func subsequencesWithMiddleMode(nums []int) int {
	left, right := map[int]int{}, map[int]int{}
	for i := 2; i < len(nums); i++ {
		right[nums[i]] += 1
	}
	for i := 0; i < 2; i++ {
		left[nums[i]] += 1
	}
	dict := map[int]bool{}
	for k := range left {
		if left[k]+right[k] >= 2 {
			dict[k] = true
		}
	}
	for k := range right {
		if left[k]+right[k] >= 2 {
			dict[k] = true
		}
	}
	var ret int
	for i := 2; i < len(nums)-2; i++ {
		leftCnt, rightCnt := i, len(nums)-i-1
		mid := nums[i]
		right[mid] -= 1
		// left 2
		ret += combination(left[mid], 2) * combination(rightCnt, 2)
		// left 1 right 2
		ret += combination(left[mid], 1) * combination(right[mid], 2) * combination(leftCnt-left[mid], 1)
		// left 1 right 1
		ret += combination(left[mid], 1) * combination(right[mid], 1) * combination(leftCnt-left[mid], 1) * combination(rightCnt-right[mid], 1)
		// left 1 right 0
		{
			tmp := combination(leftCnt-left[mid], 1) * combination(rightCnt-right[mid], 2)
			for k := range dict {
				if k == mid {
					continue
				}
				// right 2
				tmp -= combination(leftCnt-left[mid], 1) * combination(right[k], 2)
				// left 1 right 1
				tmp -= combination(left[k], 1) * combination(right[k], 1) * combination(rightCnt-right[mid]-right[k], 1)
			}
			ret += combination(left[mid], 1) * tmp
		}
		// left 0 right 2
		ret += combination(leftCnt-left[mid], 2) * combination(right[mid], 2)
		// left 0 right 1
		{
			tmp := combination(rightCnt-right[mid], 1) * combination(leftCnt-left[mid], 2)
			for k := range dict {
				if k == mid {
					continue
				}
				// right 0 left 2
				tmp -= combination(rightCnt-right[mid], 1) * combination(left[k], 2)
				// right 1 left 1
				tmp -= combination(right[k], 1) * combination(left[k], 1) * combination(leftCnt-left[mid]-left[k], 1)
			}
			ret += combination(right[mid], 1) * tmp
		}
		ret %= 1e9 + 7
		left[mid] += 1
	}
	return ret
}

func combination(a, b int) int {
	if a < b {
		return 0
	}
	if b*2 > a {
		return combination(a, a-b)
	}
	ret := 1
	for i := 0; i < b; i++ {
		ret *= (a - i)
		ret %= 1e9 + 7
		// ret /= (i + 1)
		ret *= fastPow(i+1, 1e9+7-2, 1)
		ret %= 1e9 + 7
	}
	return ret
}

func fastPow(x, k int, base int) int {
	for k != 0 {
		if k&1 != 0 {
			base *= x
			base %= 1e9 + 7
		}
		k >>= 1
		x *= x
		x %= 1e9 + 7
	}
	return base
}
