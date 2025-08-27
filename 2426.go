package main

func numberOfPairs2(nums1 []int, nums2 []int, diff int) int64 {
	diffArr := make([]int, len(nums1))
	for i := range nums1 {
		diffArr = append(diffArr, nums1[i]-nums2[i])
	}
	tmp := make([]int, 0, len(nums1))
	var helper func(l, r int) int
	helper = func(l, r int) int {
		if l == r {
			return 0
		}
		mid := (l + r) / 2
		ret := helper(l, mid) + helper(mid+1, r)
		// check num
		for i, j := l, mid+1; i <= mid || j <= r; {
			if diffArr[i] <= diffArr[j]+diff {
				ret += r - j + 1
				i++
			} else {
				j++
			}
		}
		// sort merge
		for i, j := l, mid+1; i <= mid || j <= r; {
			if i > mid || (j <= r && diffArr[j] < diffArr[i]) {
				tmp = append(tmp, diffArr[j])
				j++
			} else {
				tmp = append(tmp, diffArr[i])
				i++
			}
		}
		copy(diffArr[l:r+1], tmp)
		tmp = tmp[:0]
		return ret
	}
	return int64(helper(0, len(nums1)-1))
}
