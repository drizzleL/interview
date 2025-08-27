package main

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	cnt   map[int]int
}

func FindConstructor(nums1 []int, nums2 []int) FindSumPairs {
	m := map[int]int{}
	for _, num := range nums2 {
		m[num] += 1
	}
	ret := FindSumPairs{
		nums1: nums1,
		nums2: nums2,
		cnt:   m,
	}
	return ret
}

func (this *FindSumPairs) Add(index int, val int) {
	oldVal := this.nums2[index]
	this.cnt[oldVal] -= 1
	this.cnt[oldVal+val] += 1
	this.nums2[index] += val
}

func (this *FindSumPairs) Count(tot int) int {
	var ret int
	for _, num := range this.nums1 {
		ret += this.cnt[tot-num]
	}
	return ret
}
