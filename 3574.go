package main

func maxGCDScore(nums []int, k int) int64 {
	var ret int
	divisorCnt := func(x int, d int) int {
		var ret int
		for x%d == 0 {
			x /= d
			ret += 1
		}
		return ret
	}
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		minDC := divisorCnt(v, 2)
		cnt := 1
		ret = max(ret, v*2)
		for j := i + 1; j < len(nums); j++ {
			v = gcd(v, nums[j])
			cnt2 := divisorCnt(nums[j], 2)
			if cnt2 == minDC {
				cnt += 1
			} else if cnt2 < minDC {
				cnt = 1
				minDC = cnt2
			}
			ret = max(ret, v*(j-i+1))
			if cnt <= k {
				ret = max(ret, v*(j-i+1)*2)
			}
		}
	}
	return int64(ret)
}
