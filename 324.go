package main

import (
	"log"
	"math/rand"
)

func wiggleSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	median := getKth(nums, len(nums)/2+1)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	getIdx := func(i int) int {
		return (i*2 + 1) % (len(nums) | 1)
	}
	i, j, k := 0, 0, len(nums)-1
	for j <= k {
		idx := getIdx(j)
		switch {
		case nums[idx] > median:
			ii := getIdx(i)
			nums[ii], nums[idx] = nums[idx], nums[ii]
			i++
			j++
		case nums[idx] < median:
			kk := getIdx(k)
			nums[kk], nums[idx] = nums[idx], nums[kk]
			k--
		default:
			j++
		}
	}
}

func getKth(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	pivot := nums[0]
	for l < r {
		if nums[l+1] <= pivot {
			nums[l], nums[l+1] = nums[l+1], nums[l]
			l++
		} else {
			nums[l+1], nums[r] = nums[r], nums[l+1]
			r--
		}
	}
	if k-1 == l {
		return nums[l]
	}
	if k-1 > l {
		return getKth(nums[l+1:], k-l-1)
	}
	return getKth(nums[:l], k)
}

func wiggleSort2(nums []int) {
	N := len(nums)

	realIndex := func(i int) int {
		return (1 + 2*i) % (N | 1)
	}

	// portion must in desc order and 3-way
	portion := func(nums []int, lt, rt int) (int, int) {
		log.Println(nums)
		idx := rand.Intn(rt-lt+1) + lt
		val := nums[realIndex(idx)]
		nums[realIndex(idx)], nums[realIndex(lt)] = nums[realIndex(lt)], nums[realIndex(idx)]

		i := lt
		j := rt
		k := i
		for k <= j {
			if nums[realIndex(k)] > val {
				nums[realIndex(k)], nums[realIndex(i)] = nums[realIndex(i)], nums[realIndex(k)]
				k++
				i++
			} else if nums[realIndex(k)] < val {
				nums[realIndex(k)], nums[realIndex(j)] = nums[realIndex(j)], nums[realIndex(k)]
				j--
			} else {
				k++
			}
		}
		log.Println(nums)

		return i, j
	}

	divideHalf := func(nums []int) {
		center := len(nums) / 2
		lt := 0
		rt := len(nums) - 1
		pivotL, pivotR := rt, lt
		for pivotR < center || pivotL > center {
			pivotL, pivotR = portion(nums, lt, rt)
			if pivotL > center {
				rt = pivotL - 1
			} else {
				lt = pivotR + 1
			}
		}
	}

	shuffle := func(nums []int) {
		N := len(nums)
		for i := N - 1; i > 0; i-- {
			d := rand.Intn(i + 1)
			nums[d], nums[i] = nums[i], nums[d]
		}
	}

	shuffle(nums)
	divideHalf(nums)
	log.Println(nums)
}
