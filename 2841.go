package main

func maxSum4(nums []int, m int, k int) int64 {
	cnts := map[int]int{}
	var uniq int
	var sum, ret int
	for i := 0; i < k; i++ {
		num := nums[i]
		sum += num
		if cnts[num] == 0 {
			uniq += 1
		}
		cnts[num] += 1
	}
	if uniq >= m {
		ret = sum
	}
	for i := k; i < len(nums)-1; i++ {
		preNum := nums[i-k]
		sum -= preNum
		if cnts[preNum] == 1 {
			uniq -= 1
		}
		cnts[preNum] -= 1
		num := nums[i]
		sum += num
		if cnts[num] == 0 {
			uniq += 1
		}
		cnts[num] += 1
		if uniq >= m {
			ret = max(ret, sum)
		}
	}
	return int64(ret)
}
