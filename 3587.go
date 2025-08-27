package main

func minSwaps3(nums []int) int {
	var even, odd int
	for _, num := range nums {
		if num%2 == 0 {
			even += 1
		} else {
			odd += 1
		}
	}
	if abs(even-odd) >= 2 {
		return -1
	}
	check := func(flag int) int {
		seen := make([]bool, len(nums))
		var ret int
		var hideCnt int
		for i, j := 0, 0; i < len(nums); i++ {
			if seen[i] {
				hideCnt -= 1
				continue
			}
			if nums[i]%2 == flag {
				flag ^= 1
				j = max(j, i+1)
				continue
			}
			for nums[j]%2 != flag {
				j++
			}
			seen[j] = true
			hideCnt += 1
			j += 1
			ret += j - i - hideCnt
		}
		return ret
	}
	if even == odd {
		return min(check(0), check(1))
	}
	if even > odd {
		return check(0)
	}
	return check(1)
}
