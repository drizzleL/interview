package main

func smallestK(arr []int, k int) []int {
	return findK(arr, 0, len(arr)-1, k)
}

func findK(a []int, lo, hi, k int) []int {
	if len(a) == 0 || lo < 0 || lo >= len(a) || hi < 0 || hi >= len(a) {
		return nil
	}

	pivot := a[lo]
	i, j := lo, hi

	for i < j {
		for i < j && a[j] >= pivot {
			j--
		}
		for i < j && a[i] <= pivot {
			i++
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}

	a[lo], a[i] = a[i], a[lo]

	if i == k {
		return a[:i]
	} else if i > k {
		return findK(a, lo, i-1, k)
	} else {
		return findK(a, i+1, hi, k)
	}
}
