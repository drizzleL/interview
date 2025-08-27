package main

import "sort"

func minimumDeletions2(word string, k int) int {
	var freq [26]int
	for _, c := range word {
		freq[c-'a'] += 1
	}
	var arr []int
	for _, v := range freq {
		arr = append(arr, v)
	}
	sort.Ints(arr)
	ret := len(word)
	var presum1, presum2 int
	for i, j := 0, 0; i < len(arr); i++ {
		for j < len(arr) && arr[j]-arr[i] <= k {
			presum2 += arr[j]
			j++
		}
		ret = min(ret, presum1+len(word)-presum2-(arr[i]+k)*(len(arr)-j))
		presum1 += arr[i]
	}
	return ret
}
