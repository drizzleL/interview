package main

func numberOfSubarrays2(nums []int) int64 {
	var q [][2]int
	ret := len(nums)
	for _, num := range nums {
		for len(q) != 0 && q[len(q)-1][0] < num {
			q = q[:len(q)-1]
		}
		if len(q) == 0 || q[len(q)-1][0] > num {
			q = append(q, [2]int{num, 1})
			continue
		}
		ret += q[len(q)-1][1]
		q[len(q)-1][1] += 1
	}
	return int64(ret)
}
