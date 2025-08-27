package main

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	divide := func(nums []int) (neg []int, pos []int) {
		for _, num := range nums {
			if num < 0 {
				neg = append(neg, num)
			} else {
				pos = append(pos, num)
			}
		}
		return
	}
	neg1, pos1 := divide(nums1)
	neg2, pos2 := divide(nums2)
	rev := func(arr []int) []int {
		ret := make([]int, len(arr))
		for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
			ret[j] = arr[i]
		}
		return ret
	}
	pos1Rev, pos2Rev := rev(pos1), rev(pos2)
	neg1Rev, neg2Rev := rev(neg1), rev(neg2)
	arr := []int{nums1[0] * nums2[0], nums1[0] * nums2[len(nums2)-1], nums1[len(nums1)-1] * nums2[len(nums2)-1], nums1[len(nums1)-1] * nums2[0]}
	l, r := arr[0], arr[0]
	for _, v := range arr {
		l = min(l, v)
		r = max(r, v)
	}
	getCnt := func(x int, a, b []int) int {
		var ret int
		for i, j := 0, len(b)-1; i < len(a); i++ {
			for j >= 0 && a[i]*b[j] > x {
				j--
			}
			ret += j + 1
		}
		return ret
	}
	for l < r {
		mid := l + (r-l-1)/2
		var cnt int
		if mid >= 0 { // only consider neg
			cnt += len(pos1) * len(neg2)
			cnt += len(pos2) * len(neg1)
			cnt += getCnt(mid, neg1Rev, neg2Rev)
			cnt += getCnt(mid, pos1, pos2)
		} else {
			cnt += getCnt(mid, pos2Rev, neg1)
			cnt += getCnt(mid, pos1Rev, neg2)
		}
		if cnt >= int(k) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return int64(l)
}
